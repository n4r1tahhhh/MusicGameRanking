package model

import (
	"gorm.io/gorm"
)

// Competition 大会
type Competition struct {
	gorm.Model
	Name string `json:"name"`
}
