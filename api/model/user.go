package model

import (
	"gorm.io/gorm"
)

// User 会員
type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash []byte `json:"password"`
}
