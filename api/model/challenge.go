package model

import (
	"gorm.io/gorm"
)

// Challenge 課題
type Challenge struct {
	gorm.Model
	Name string `json:"name"`
}
