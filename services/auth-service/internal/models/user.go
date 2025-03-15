package models

import (
	"time"
)

type User struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	Username   string    `json:"username" gorm:"unique;not null"`
	Password   string    `json:"password" gorm:"not null"`
	Email      string    `json:"email" gorm:"unique;not null"`
	UserRole   string    `json:"user_role"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
