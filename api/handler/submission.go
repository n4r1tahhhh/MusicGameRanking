package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/n4r1tahhhh/MusicGameRanking/model"
)

// CreateSubmission スコアを作成
func (h *Handler) CreateSubmission(c echo.Context) error {
	score, err := strconv.Atoi(c.FormValue("score"))
	if err != nil {
		return err
	}

	// Validate

	// Bind
	submission := &model.Submission{
		Score: score,
	}

	// Save submission
	if result := h.DB.Model(&model.Submission{}).Create(submission); result.Error != nil {
		return result.Error
	}
	return c.String(http.StatusCreated, "Created!")
}

// GetSubmissions スコア一覧の取得
func (h *Handler) GetSubmissions(c echo.Context) error {
	var submissions []model.Submission
	if result := h.DB.Find(&submissions); result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, submissions)
}
