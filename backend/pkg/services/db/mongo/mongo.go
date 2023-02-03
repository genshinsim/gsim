package mongo

import (
	"context"
	"os"
	"strconv"

	"github.com/genshinsim/gcsim/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Config struct {
	URL        string
	Database   string
	Collection string
	QueryView  string
	Username   string
	Password   string
}

type Server struct {
	cfg          Config
	client       *mongo.Client
	Log          *zap.SugaredLogger
	maxPageLimit int64
}

func NewServer(cfg Config, cust ...func(*Server) error) (*Server, error) {
	s := &Server{
		cfg:          cfg,
		maxPageLimit: 100,
	}

	limitStr := os.Getenv("MONGO_STORE_MAX_LIMIT")
	if limit, err := strconv.ParseInt(limitStr, 10, 64); err == nil && limit > 0 {
		s.maxPageLimit = limit
	}

	for _, f := range cust {
		err := f(s)
		if err != nil {
			return nil, err
		}
	}

	if s.Log == nil {
		logger, err := zap.NewProduction()
		if err != nil {
			return nil, err
		}
		sugar := logger.Sugar()
		sugar.Debugw("logger initiated")

		s.Log = sugar
	}

	credential := options.Credential{
		Username: cfg.Username,
		Password: cfg.Password,
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.URL).SetAuth(credential))
	if err != nil {
		return nil, err
	}

	//check connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		s.Log.Errorw("mongodb ping failed", "err", err)
		return nil, err
	}

	s.Log.Info("mongodb connected sucessfully")

	s.client = client

	return s, nil
}

func (s *Server) Create(ctx context.Context, entry *model.DBEntry) (string, error) {
	s.Log.Infow("mongodb: create request", "entry", entry.String())

	col := s.client.Database(s.cfg.Database).Collection(s.cfg.Collection)

	res, err := col.InsertOne(ctx, entry)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			s.Log.Infow("mongodb: create failed - duplicated id", "key", entry.Key, "err", err)
			return "", status.Error(codes.InvalidArgument, "duplicated id")
		}
		s.Log.Infow("mongodb: create failed - unexpected error", "key", entry.Key, "err", err)
		return "", status.Error(codes.Internal, "internal server error")
	}

	s.Log.Infow("mongodb: create successful", "id", res.InsertedID)

	return entry.Key, nil
}

type paginate struct {
	limit int64
	page  int64
}

func newPaginate(limit, page int64) *paginate {
	return &paginate{
		limit: int64(limit),
		page:  int64(page),
	}
}

func (p *paginate) opts() *options.FindOptions {
	l := p.limit
	skip := p.page*p.limit - p.limit
	fOpt := options.FindOptions{Limit: &l, Skip: &skip}

	return &fOpt
}

func (s *Server) parseLimit(opt *model.DBQueryOpt) int64 {
	limit := opt.GetLimit()
	if limit < 0 {
		return 0
	}
	if limit > s.maxPageLimit {
		return s.maxPageLimit
	}
	return limit
}

func (s *Server) Get(ctx context.Context, opt *model.DBQueryOpt) ([]*model.DBEntry, error) {
	s.Log.Infow("mongodb: get request", "query", opt)

	col := s.client.Database(s.cfg.Database).Collection(s.cfg.QueryView)

	cursor, err := col.Find(ctx, opt.GetQuery().AsMap(), newPaginate(s.parseLimit(opt), opt.GetSkip()).opts())
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, "no records found")
		}
		return nil, status.Error(codes.Internal, "unexpected server error")
	}

	var res []model.DBEntry

	if err = cursor.All(ctx, &res); err != nil {
		return nil, status.Error(codes.Internal, "unexpected server error")
	}

	if len(res) == 0 {
		s.Log.Infow("mongodb: get request done; no results")
		return nil, nil
	}

	var result []*model.DBEntry

	for i := 0; i < len(res); i++ {
		r := &res[i]
		cursor.Decode(r)
		result = append(result, r)
	}

	s.Log.Infow("mongodb: get request done", "count", len(result))

	return result, nil
}
