package main

import (
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/warrenb95/railway-signals/internal/adapters/http"
	"github.com/warrenb95/railway-signals/internal/adapters/repository"
	"github.com/warrenb95/railway-signals/internal/application"
)

// Database connection setup
func connectDB() *pg.DB {
	opts := &pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "password",
		Database: "railway_db",
	}
	return pg.Connect(opts)
}

func main() {
	db := connectDB()
	defer db.Close()
	logger := logrus.New()

	repo, err := repository.NewRepository(db, logger)
	if err != nil {
		logger.WithError(err).Fatal("Creating new repository")
	}

	svr := &application.Service{
		Logger:       logger,
		SignalStore:  repo,
		TrackStore:   repo,
		MileageStore: repo,
	}

	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define API routes
	// TODO: api groups?
	e.GET("/v1/signals", http.ListSignalHandler(svr))
	e.GET("/v1/signals/:id", http.GetSignalHandler(svr))
	e.POST("/v1/signals", http.CreateSignalHandler(svr))
	e.PUT(("/v1/signals/:id"), http.UpdateSignalHandler(svr))
	e.DELETE(("/v1/signals/:id"), http.DeleteSignalHandler(svr))

	e.GET("/v1/tracks", http.ListTrackHandler(svr))
	e.GET("/v1/tracks/:id", http.GetTrackHandler(svr))
	e.POST("/v1/tracks", http.CreateTrackHandler(svr))
	e.PUT("/v1/tracks/:id", http.UpdateTrackHandler(svr))
	e.DELETE("/v1/tracks/:id", http.DeleteTrackHandler(svr))

	e.GET("/v1/signals/:id/tracks", http.GetSignalTracks(svr))

	e.Logger.Fatal(e.Start(":8080"))
}
