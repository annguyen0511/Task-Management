package models

import "time"

type Role struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Permissions []string  `json:"permissions"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
