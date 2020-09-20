package handler

import (
	"gorm.io/gorm"
)

// Handler からDB操作するための準備
type (
	Handler struct {
		DB *gorm.DB
	}
)
