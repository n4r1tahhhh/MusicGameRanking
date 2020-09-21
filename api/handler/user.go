package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/n4r1tahhhh/MusicGameRanking/model"
	"golang.org/x/crypto/bcrypt"
)

// Signup 新規登録
func (h *Handler) Signup(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Validate

	// パスワードのハッシュ化
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Fatal(err)
	}

	// Bind
	u := &model.User{
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
	}

	// Save user
	if err := h.DB.Create(u); err != nil {
		log.Fatal(err)
	}
	res := fmt.Sprintf(`{"name":"%s", "email":"%s"}`, u.Name, u.Email)
	return c.JSON(http.StatusCreated, res)
}

// GetUsers ユーザ一覧の取得
func (h *Handler) GetUsers(c echo.Context) error {
	results := map[string]interface{}{}
	h.DB.Model(&model.User{}).Select("name", "email").Find(&results)
	return c.JSON(http.StatusOK, results)
}
