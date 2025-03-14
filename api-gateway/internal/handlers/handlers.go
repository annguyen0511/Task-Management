package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/annguyen0511/Task-Management/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Handler chứa các client gRPC để gọi microservices
type Handler struct {
	AuthClient proto.AuthServiceClient
	TaskClient proto.TaskServiceClient
}

func NewHandler(authAddr, taskAddr string) (*Handler, error) {

	authConn, err := grpc.Dial(authAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	taskConn, err := grpc.Dial(taskAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Handler{
		AuthClient: proto.NewAuthServiceClient(authConn),
		TaskClient: proto.NewTaskServiceClient(taskConn),
	}, nil
}

func (h *Handler) Login(c *gin.Context) {
	var req proto.LoginRequest
	// BindJSON binds the passed struct pointer using the request body as JSON.
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.AuthClient.Login(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to auth service"})
		return
	}

	if resp.Error != "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": resp.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": resp.Token})
}

func (h *Handler) Register(c *gin.Context) {
	var req proto.RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.AuthClient.Register(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to auth service"})
		return
	}

	if resp.Error != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (h *Handler) GetTasks(c *gin.Context) {
	userID := c.Query("user_id") // Get user_id from query param string
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}
	req := &proto.GetTasksRequest{UserId: int32(userIDInt)}
	resp, err := h.TaskClient.GetTasks(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to task service"})
		return
	}

	if resp.Error != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": resp.Tasks})
}

func (h *Handler) GetTask(c *gin.Context) {
	userID := c.Query("user_id")
	taskID := c.Param("id")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}
	taskIDInt, err := strconv.Atoi(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task_id"})
		return
	}
	req := &proto.GetTaskRequest{UserId: int32(userIDInt), Id: int32(taskIDInt)}
	resp, err := h.TaskClient.GetTask(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to task service"})
		return
	}

	if resp.Error != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"title": resp.Title, "description": resp.Description})

}

func (h *Handler) CreateTask(c *gin.Context) {
	var req proto.CreateTaskRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.TaskClient.CreateTask(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to task service"})
		return
	}

	if resp.Error != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully", "task_id": resp.Id})
}
