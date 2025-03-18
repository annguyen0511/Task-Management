package handlers

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/annguyen0511/Task-Management/proto"
	"github.com/annguyen0511/Task-Management/services/auth-service/config"
	"github.com/annguyen0511/Task-Management/services/auth-service/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type AuthServer struct {
	proto.UnimplementedAuthServiceServer
	DB     *gorm.DB
	Config config.Config
}

func NewAuthServer(db *gorm.DB, cfg config.Config) *AuthServer {
	return &AuthServer{
		DB:     db,
		Config: cfg,
	}
}

func (s *AuthServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	var user models.User
	hashedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password)))
	if err := s.DB.Where("email = ? AND password = ?", req.Email, hashedPassword).First(&user).Error; err != nil {
		return &proto.LoginResponse{Error: "Invalid email or password"}, nil
	}

	token, err := generateJWT(user.ID, s.Config.JWTSecret)
	if err != nil {
		return &proto.LoginResponse{Error: "Failed to generate token"}, nil
	}
	return &proto.LoginResponse{Token: token}, nil
}

func (s *AuthServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	hashedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password)))
	user := models.User{
		Email:    req.Email,
		Password: hashedPassword,
	}
	if err := s.DB.Create(&user).Error; err != nil {
		return &proto.RegisterResponse{Error: "Failed to create user"}, nil
	}
	token, err := generateJWT(user.ID, s.Config.JWTSecret)
	if err != nil {
		return &proto.RegisterResponse{Error: "Failed to generate token"}, nil
	}
	return &proto.RegisterResponse{Msg: "Created user successfully", Token: token}, nil
}

func generateJWT(userID uint, secret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // JWT will expire in 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
