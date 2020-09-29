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

func TestSubmission(t *testing.T) {
	// Setup(CreateSubmission)
	f := make(url.Values)
	f.Set("score", "100")
	req := httptest.NewRequest(http.MethodPost, "/submissions", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := E.NewContext(req, rec)

	// Assertions(Createsubmission)
	if assert.NoError(t, H.CreateSubmission(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "Created!", rec.Body.String())
	}

	// Setup(GetSubmissions)
	req = httptest.NewRequest(http.MethodGet, "/submissions", nil)
	rec = httptest.NewRecorder()
	c = E.NewContext(req, rec)

	// Assertions(GetSubmissions)
	if assert.NoError(t, H.GetSubmissions(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "100")
	}
}
