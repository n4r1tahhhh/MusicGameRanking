package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	// Setup(Signup)
	f := make(url.Values)
	f.Set("name", "Jon")
	f.Set("email", "jon@example.com")
	f.Set("password", "password")
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := E.NewContext(req, rec)

	// Assertions(Signup)
	if assert.NoError(t, H.Signup(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "Created!", rec.Body.String())
	}

	// Setup(GetUsers)
	req = httptest.NewRequest(http.MethodGet, "/users", nil)
	rec = httptest.NewRecorder()
	c = E.NewContext(req, rec)

	// Assertions(GetUsers)
	if assert.NoError(t, H.GetUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "[{\"Name\":\"Jon\",\"Email\":\"jon@example.com\"}]\n", rec.Body.String())
	}
}
