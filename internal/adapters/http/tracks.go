package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/warrenb95/railway-signals/internal/application"
	"github.com/warrenb95/railway-signals/internal/domain"
)

func CreateTrackHandler(s *application.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var track domain.Track
		if err := c.Bind(&track); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		if err := s.CreateTrack(c.Request().Context(), &track); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create track"})
		}

		return c.JSON(http.StatusCreated, track)
	}
}

func GetTrackHandler(s *application.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		trackIDStr := c.Param("id")
		if trackIDStr == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Empty track ID"})
		}

		trackID, err := strconv.Atoi(trackIDStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid track ID"})
		}

		track, err := s.GetTrack(c.Request().Context(), trackID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create track"})
		}

		return c.JSON(http.StatusOK, track)
	}
}

func ListTrackHandler(s *application.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var page int
		limit := int(100) // Default limit TODO: should be configureable?
		if pageStr := c.Param("page"); pageStr != "" {
			p, err := strconv.Atoi(pageStr)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid pagination page value"})
			}
			page = p
		}

		if limitStr := c.Param("limit"); limitStr != "" {
			l, err := strconv.Atoi(limitStr)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid pagination limit value"})
			}
			limit = l
		}

		tracks, count, err := s.ListTracks(c.Request().Context(), limit, page)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to list tracks"})
		}

		var nextPage int
		if page*limit < count {
			nextPage = page + 1
		}

		return c.JSON(http.StatusOK, map[string]any{
			"tracks":    tracks,
			"next_page": nextPage,
		})
	}
}

func UpdateTrackHandler(s *application.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var track domain.Track
		if err := c.Bind(&track); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		if err := s.UpdateTrack(c.Request().Context(), &track); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update track"})
		}

		return c.JSON(http.StatusOK, track)
	}
}

func DeleteTrackHandler(s *application.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		trackIDstr := c.Param("id")
		if trackIDstr == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Empty track ID"})
		}

		trackID, err := strconv.Atoi(trackIDstr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid track ID"})
		}

		err = s.DeleteTrack(c.Request().Context(), trackID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete track"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Deleted successfully"})
	}
}
