package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/n4r1tahhhh/MusicGameRanking/model"
)

func (h *Handler) Signup(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Validate

	// Bind
	u := &model.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	// Save user
	h.DB.Create(u)
	return c.JSON(http.StatusCreated, u)
}
