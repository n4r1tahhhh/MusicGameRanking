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

func TestCompetition(t *testing.T) {
	// Setup(CreateCompetitions)
	f := make(url.Values)
	f.Set("name", "Taiko no Tatsujin competition")
	req := httptest.NewRequest(http.MethodPost, "/competitions", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := E.NewContext(req, rec)

	// Assertions(CreateCompetitions)
	if assert.NoError(t, H.CreateCompetitions(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "Created!", rec.Body.String())
	}

	// Setup(GetCompetitions)
	req = httptest.NewRequest(http.MethodGet, "/competitions", nil)
	rec = httptest.NewRecorder()
	c = E.NewContext(req, rec)

	// Assertions(GetCompetitions)
	if assert.NoError(t, H.GetCompetitions(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Taiko no Tatsujin")
	}
}
