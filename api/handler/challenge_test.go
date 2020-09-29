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

func TestChallenge(t *testing.T) {
	// Setup(CreateChallenge)
	f := make(url.Values)
	f.Set("name", "music")
	req := httptest.NewRequest(http.MethodPost, "/challenges", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := E.NewContext(req, rec)

	// Assertions(CreateChallenge)
	if assert.NoError(t, H.CreateChallenge(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "Created!", rec.Body.String())
	}

	// Setup(GetChallenges)
	req = httptest.NewRequest(http.MethodGet, "/challenges", nil)
	rec = httptest.NewRecorder()
	c = E.NewContext(req, rec)

	// Assertions(GetChallenges)
	if assert.NoError(t, H.GetChallenges(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "music")
	}
}
