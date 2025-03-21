package main

import (
	"log"
	"net"

	"github.com/annguyen0511/Task-Management/proto"
	"github.com/annguyen0511/Task-Management/services/task-service/config"
	"github.com/annguyen0511/Task-Management/services/task-service/internal/db"
	"github.com/annguyen0511/Task-Management/services/task-service/internal/handlers"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbConn, err := db.InitDB(*cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close(dbConn)

	lis, err := net.Listen("tcp", ":"+cfg.ServerPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterTaskServiceServer(s, handlers.NewTaskServer(dbConn))
	log.Printf("Task service is running on port %s", cfg.ServerPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
