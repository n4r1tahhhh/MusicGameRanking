package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/n4r1tahhhh/MusicGameRanking/handler"
	"github.com/n4r1tahhhh/MusicGameRanking/model"
)

func main() {
	e := echo.New()

	// Database connection
	db, err := handler.ConnectDB(false) // isTest = false
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Migration
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Competition{})
	db.AutoMigrate(&model.Challenge{})
	db.AutoMigrate(&model.Submission{})

	// Initialize handler
	h := &handler.Handler{DB: db}

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// user
	e.GET("/users", h.GetUsers)
	e.POST("/users", h.Signup)
	// competition
	e.GET("/competitions", h.GetCompetitions)
	e.POST("/competitions", h.CreateCompetition)
	// challenge
	e.GET("/challenges", h.GetChallenges)
	e.POST("/challenges", h.CreateChallenge)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
