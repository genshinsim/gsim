package main

import (
	"log"
	"net"
	"os"

	"github.com/genshinsim/gcsim/backend/pkg/mongo"
	"github.com/genshinsim/gcsim/backend/pkg/services/share"
	"github.com/genshinsim/gcsim/backend/pkg/services/submission"
	"google.golang.org/grpc"
)

func main() {
	mongoCfg := mongo.Config{
		URL:        os.Getenv("MONGODB_URL"),
		Database:   os.Getenv("MONGODB_DATABASE"),
		Collection: os.Getenv("MONGODB_COLLECTION"),
		QueryView:  os.Getenv("MONGODB_QUERY_VIEW"),
		Username:   os.Getenv("MONGODB_USERNAME"),
		Password:   os.Getenv("MONOGDB_PASSWORD"),
	}
	log.Println(os.Getenv("MONGODB_URL"))
	log.Printf("Cfg: %v\n", mongoCfg)
	dbStore, err := mongo.NewServer(mongoCfg)
	if err != nil {
		panic(err)
	}
	shareStore, err := share.NewClient(share.ClientCfg{
		Addr: os.Getenv("SHARE_STORE_URL"),
	})

	if err != nil {
		panic(err)
	}

	server, err := submission.NewServer(submission.Config{
		DBStore:    dbStore,
		ShareStore: shareStore,
	})

	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen on port 3000: %v", err)
	}

	s := grpc.NewServer()
	submission.RegisterSubmissionStoreServer(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
