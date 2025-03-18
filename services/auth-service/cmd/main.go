package main

import (
	"log"
	"net"

	"github.com/annguyen0511/Task-Management/proto"
	"github.com/annguyen0511/Task-Management/services/auth-service/config"
	"github.com/annguyen0511/Task-Management/services/auth-service/internal/db"
	"github.com/annguyen0511/Task-Management/services/auth-service/internal/handlers"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	dbConn, err := db.InitDB(*cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close(dbConn)

	lis, err := net.Listen("tcp", ":"+cfg.ServerPort)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	proto.RegisterAuthServiceServer(s, handlers.NewAuthServer(dbConn, *cfg))
	log.Printf("Starting server on port %s", cfg.ServerPort)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}

}
