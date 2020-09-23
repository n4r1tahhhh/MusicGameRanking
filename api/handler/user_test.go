package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/n4r1tahhhh/MusicGameRanking/model"
	"github.com/stretchr/testify/assert"
)

var e *echo.Echo
var h *Handler

func TestMain(m *testing.M) {
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

	code := m.Run()

	db.Exec("DROP TABLE users")
	os.Exit(code)
}

func TestUser(t *testing.T) {
	// Setup(Signup)
	f := make(url.Values)
	f.Set("name", "Jon")
	f.Set("email", "jon@example.com")
	f.Set("password", "password")
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions(Signup)
	if assert.NoError(t, h.Signup(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "Created!", rec.Body.String())
	}

	// Setup(GetUsers)
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	// Assertions(GetUsers)
	if assert.NoError(t, h.GetUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "[{\"Name\":\"Jon\",\"Email\":\"jon@example.com\"}]\n", rec.Body.String())
	}
}
