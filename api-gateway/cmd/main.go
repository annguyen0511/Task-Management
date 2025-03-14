package main

import (
	"log"

	"github.com/annguyen0511/Task-Management/api-gateway/config"
	"github.com/annguyen0511/Task-Management/api-gateway/internal/handlers"
	"github.com/annguyen0511/Task-Management/api-gateway/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	h, err := handlers.NewHandler(cfg.AuthServiceAddr, cfg.TasksServiceAddr)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	routes.SetupRoutes(r, h)

	log.Printf("API Gateway starting server on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		panic(err)
	}
}
