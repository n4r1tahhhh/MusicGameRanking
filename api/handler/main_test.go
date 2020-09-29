package handler

import (
	"os"
	"testing"

	"github.com/labstack/echo"
	"github.com/n4r1tahhhh/MusicGameRanking/model"
)

var E *echo.Echo
var H *Handler

func TestMain(m *testing.M) {
	E = echo.New()

	// Database connection
	db, err := ConnectDB(true) // isTest = true
	if err != nil {
		E.Logger.Fatal(err)
	}

	// Migration
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Competition{})
	db.AutoMigrate(&model.Challenge{})

	// Initialize handler
	H = &Handler{DB: db}

	code := m.Run()

	db.Exec("DROP TABLE users")
	db.Exec("DROP TABLE competitions")
	os.Exit(code)
}
