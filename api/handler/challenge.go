package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/n4r1tahhhh/MusicGameRanking/model"
)

// CreateChallenge 課題を作成
func (h *Handler) CreateChallenge(c echo.Context) error {
	name := c.FormValue("name")

	// Validate

	// Bind
	challenge := &model.Challenge{
		Name: name,
	}

	// Save Challenge
	if result := h.DB.Model(&model.Challenge{}).Create(challenge); result.Error != nil {
		return result.Error
	}
	return c.String(http.StatusCreated, "Created!")
}

// GetChallenges 課題一覧の取得
func (h *Handler) GetChallenges(c echo.Context) error {
	var challenges []model.Challenge
	if result := h.DB.Find(&challenges); result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, challenges)
}
