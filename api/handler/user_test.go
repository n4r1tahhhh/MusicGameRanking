package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/n4r1tahhhh/MusicGameRanking/model"
	"github.com/stretchr/testify/assert"
)

var e *echo.Echo
var h *Handler
var userJSON = `{"name":"Jon", "email":"Jon@example.com", "password":"password"}`

func TestMain(t *testing.T) {
	e = echo.New()

	// Database connection
	db, err := ConnectDB(true) // isTest = true
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Migration
	db.AutoMigrate(&model.User{})

	// Initialize handler
	h = &Handler{DB: db}

	t.Run("User", TestSignup)
}

func TestSignup(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.Signup(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "Created!", rec.Body.String())
	}
}
