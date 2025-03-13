package models

import "time"

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	UserRole   string    `json:"user_role"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
