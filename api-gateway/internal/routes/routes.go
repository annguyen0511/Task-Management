package routes

import (
	"github.com/annguyen0511/Task-Management/api-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handlers.Handler) {
	r.POST("/auth/login", h.Login)
	r.POST("/auth/register", h.Register)
	r.GET("/tasks", h.GetTasks)
	r.POST("/tasks", h.CreateTask)
	r.GET("/tasks/:id", h.GetTask)
}
