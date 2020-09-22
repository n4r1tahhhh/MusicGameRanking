package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/n4r1tahhhh/MusicGameRanking/model"
	"golang.org/x/crypto/bcrypt"
)

// Result APIとして実際に返すレスポンス(パスワードなどは返さない)
type Result struct {
	Name  string
	Email string
}

// Signup 新規登録
func (h *Handler) Signup(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Validate

	// パスワードのハッシュ化
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	// Bind
	u := &model.User{
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
	}

	// Save user
	if result := h.DB.Model(&model.User{}).Create(u); result.Error != nil {
		return result.Error
	}
	return c.String(http.StatusCreated, "Created!")
}

// GetUsers ユーザ一覧の取得
func (h *Handler) GetUsers(c echo.Context) error {
	// ユーザ一覧をDBからフェッチ
	var users []model.User
	if result := h.DB.Find(&users); result.Error != nil {
		return result.Error
	}

	// 名前とEmailだけ返す
	var results []Result
	for _, v := range users {
		results = append(results, Result{Name: v.Name, Email: v.Email})
	}
	return c.JSON(http.StatusOK, results)
}
