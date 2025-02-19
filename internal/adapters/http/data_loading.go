package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/warrenb95/railway-signals/internal/application"
	"github.com/warrenb95/railway-signals/internal/domain"
)

// LoadJSON will load in a list of tracks and their associated signals.
func LoadJSON(s *application.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input []domain.TrackSignals
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		if err := s.LoadTrackSignals(c.Request().Context(), input); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.NoContent(http.StatusCreated)
	}
}
