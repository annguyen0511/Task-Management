package handlers

import (
	"context"

	"github.com/annguyen0511/Task-Management/proto"
	"github.com/annguyen0511/Task-Management/services/task-service/internal/models"
	"gorm.io/gorm"
)

type TaskServer struct {
	proto.UnimplementedTaskServiceServer
	DB *gorm.DB
}

func NewTaskServer(db *gorm.DB) *TaskServer {
	return &TaskServer{
		DB: db,
	}
}

func (s *TaskServer) CreateTask(ctx context.Context, req *proto.CreateTaskRequest) (*proto.CreateTaskResponse, error) {
	task := models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		DueDate:     req.DueDate,
	}
	if err := s.DB.Create(&task).Error; err != nil {
		return &proto.CreateTaskResponse{Error: "Failed to create task"}, nil
	}
	return &proto.CreateTaskResponse{Id: task.ID}, nil
}

func (s *TaskServer) GetTasks(ctx context.Context, req *proto.GetTasksRequest) (*proto.GetTasksResponse, error) {
	var tasks []models.Task
	if err := s.DB.Where("user_id = ?", req.UserId).Find(&tasks).Error; err != nil {
		return &proto.GetTasksResponse{Error: "Failed to get tasks"}, nil
	}

	protoTasks := make([]*proto.Task, len(tasks))
	for i, task := range tasks {
		protoTasks[i] = &proto.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
		}
	}
	return &proto.GetTasksResponse{Tasks: protoTasks}, nil
}

func (s *TaskServer) GetTask(ctx context.Context, req *proto.GetTaskRequest) (*proto.GetTaskResponse, error) {
	var task models.Task
	if err := s.DB.Where("id = ? AND user_id = ?", req.Id, req.UserId).First(&task).Error; err != nil {
		return &proto.GetTaskResponse{Error: "Task not found"}, nil
	}
	return &proto.GetTaskResponse{
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		DueDate:     task.DueDate,
	}, nil
}

func (s *TaskServer) UpdateTask(ctx context.Context, req *proto.UpdateTaskRequest) (*proto.UpdateTaskResponse, error) {
	var task models.Task
	if err := s.DB.Where("id = ? AND user_id = ?", req.Id, req.UserId).First(&task).Error; err != nil {
		return &proto.UpdateTaskResponse{Error: "Task not found"}, nil
	}
	task.Title = req.Title
	task.Description = req.Description
	task.Status = req.Status
	task.DueDate = req.DueDate
	if err := s.DB.Save(&task).Error; err != nil {
		return &proto.UpdateTaskResponse{Error: "Failed to update task"}, nil
	}
	return &proto.UpdateTaskResponse{Id: task.ID}, nil
}

func (s *TaskServer) DeleteTask(ctx context.Context, req *proto.DeleteTaskRequest) (*proto.DeleteTaskResponse, error) {
	var task models.Task
	if err := s.DB.Where("id = ? AND user_id = ?", req.Id, req.UserId).First(&task).Error; err != nil {
		return &proto.DeleteTaskResponse{Error: "Task not found"}, nil
	}

	if err := s.DB.Delete(&task).Error; err != nil {
		return &proto.DeleteTaskResponse{Error: "Failed to delete task"}, nil
	}
	return &proto.DeleteTaskResponse{Id: task.ID}, nil
}
