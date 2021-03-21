package model

import (
	"time"
)

type User struct {
	Base
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     int8   `json:"role"`
}

type Base struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
