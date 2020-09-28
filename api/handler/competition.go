package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/n4r1tahhhh/MusicGameRanking/model"
)

// CreateCompetitions 大会を作成
func (h *Handler) CreateCompetitions(c echo.Context) error {
	name := c.FormValue("name")

	// Validate

	// Bind
	competition := &model.Competition{
		Name: name,
	}

	// Save competition
	if result := h.DB.Model(&model.Competition{}).Create(competition); result.Error != nil {
		return result.Error
	}
	return c.String(http.StatusCreated, "Created!")
}

// GetCompetitions 大会一覧の取得
func (h *Handler) GetCompetitions(c echo.Context) error {
	var competitions []model.Competition
	if result := h.DB.Find(&competitions); result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, competitions)
}
