package model

import (
	"gorm.io/gorm"
)

// Submission スコア
type Submission struct {
	gorm.Model
	Score int `json:"score"`
}
