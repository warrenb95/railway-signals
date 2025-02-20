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

// TODO: move this to repository layer
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

	repo, err := repository.NewPostgresRepository(db, logger)
	if err != nil {
		logger.WithError(err).Fatal("Creating new repository")
	}

	s := &application.Service{
		Logger:       logger,
		SignalStore:  repo,
		TrackStore:   repo,
		MileageStore: repo,
	}

	e := echo.New()

	// middleware
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			logger.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	}))
	e.Use(middleware.Recover())

	// Define API routes
	// TODO: api groups?
	e.GET("/v1/signals", http.ListSignalHandler(s))
	e.GET("/v1/signals/:id", http.GetSignalHandler(s))
	e.POST("/v1/signals", http.CreateSignalHandler(s))
	e.PUT(("/v1/signals/:id"), http.UpdateSignalHandler(s))
	e.DELETE(("/v1/signals/:id"), http.DeleteSignalHandler(s))

	e.GET("/v1/tracks", http.ListTrackHandler(s))
	e.GET("/v1/tracks/:id", http.GetTrackHandler(s))
	e.POST("/v1/tracks", http.CreateTrackHandler(s))
	e.PUT("/v1/tracks/:id", http.UpdateTrackHandler(s))
	e.DELETE("/v1/tracks/:id", http.DeleteTrackHandler(s))

	e.GET("/v1/signals/:id/tracks", http.GetSignalTracks(s))

	e.POST("/v1/tracks/load", http.LoadJSON(s))

	e.Logger.Fatal(e.Start(":8080"))
}
