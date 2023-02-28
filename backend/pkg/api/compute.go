package api

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/genshinsim/gcsim/pkg/model"
)

type QueueService interface {
	Get(context.Context) (*model.ComputeWork, error)
	Complete(context.Context, string) error
}

func (s *Server) computeAPIKeyCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-GCSIM-COMPUTE-API-KEY")
		s.Log.Infow("compute key check", "key", key)
		if key == "" {
			s.Log.Infow("compute key missing", "header", r.Header)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		if key != s.cfg.ComputeAPIKey {
			s.Log.Infow("compute key missing", "header", r.Header)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// callback endpoint for compute instance to submit result
func (s *Server) computeCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Log.Info("submitting compute work")

		data, err := io.ReadAll(r.Body)
		if err != nil {
			s.Log.Errorw("unexpected error reading request body", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !json.Valid(data) {
			s.Log.Info("request is not valid json")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		str := r.Header.Get("X-GCSIM-SIGNED-HASH")
		if str == "" {
			s.Log.Infow("compute callback request failed - no hash received", "header", r.Header)
			http.Error(w, "Missing Signature", http.StatusBadRequest)
			return
		}

		err = s.validateSigning(data, str)
		if err != nil {
			s.Log.Infow("create share request - validation failed", "err", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		src := r.Header.Get("X-GCSIM-COMPUTE-SRC")
		if src == "" {
			s.Log.Infow("compute callback request failed - no id received", "header", r.Header)
			http.Error(w, "Missing Src", http.StatusBadRequest)
			return
		}

		res := &model.SimulationResult{}
		err = res.UnmarshalJson(data)
		if err != nil {
			s.Log.Infow("compute callback request - unmarshall failed", "err", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		//check hash matches current
		if res.GetSimVersion() != s.cfg.CurrentHash {
			s.Log.Infow("compute callback request - invalid hash", "expected", s.cfg.CurrentHash, "got", res.GetSimVersion())
			http.Error(w, "Bad Hash", http.StatusForbidden)
			return
		}

		//check to make sure this one is actually queued
		id := r.Header.Get("X-GCSIM-COMPUTE-ID")
		if id == "" {
			s.Log.Infow("compute callback request failed - no id received", "header", r.Header)
			http.Error(w, "Missing ID", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		err = s.cfg.QueueService.Complete(ctx, id)
		if err != nil {
			s.Log.Infow("submitted work could not be completed", "err", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		//TODO: proto should be handling the constants here
		workSrc := model.ComputeWorkSource_value[src]
		switch model.ComputeWorkSource(workSrc) {
		case model.ComputeWorkSource_DBWork:
			err = s.cfg.DBStore.Update(ctx, id, res)
		case model.ComputeWorkSource_SubmissionWork:
			err = s.cfg.SubmissionStore.Complete(ctx, id, res)
		}

		if err != nil {
			s.Log.Infow("submitted work could not be completed - err adding to db", "err", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

	}
}

func (s *Server) getWork() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Log.Infow("get compute work request received")
		work, err := s.cfg.QueueService.Get(r.Context())
		if err != nil {
			s.Log.Infow("error getting work from queue", "err", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		if work == nil {
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte("no work"))
			return
		}

		data, err := work.MarshalJson()
		if err != nil {
			s.Log.Warnw("error get compute work - cannot marshal result", "err", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		writer, err := gzip.NewWriterLevel(w, gzip.BestCompression)
		if err != nil {
			s.Log.Warnw("error get comptue work - cannot write gzip result", "err", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		defer writer.Close()

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		w.WriteHeader(http.StatusOK)
		writer.Write(data)
	}
}
