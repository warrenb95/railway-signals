package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/warrenb95/railway-signals/internal/application"
	"github.com/warrenb95/railway-signals/internal/domain"
)

func CreateSignalHandler(s *application.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var signal domain.Signal
		if err := c.Bind(&signal); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		if err := s.CreateSignal(c.Request().Context(), &signal); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusCreated, signal)
	}
}

func GetSignalHandler(s *application.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		signalIDstr := c.Param("id")
		if signalIDstr == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Empty signal ID"})
		}

		signalID, err := strconv.Atoi(signalIDstr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid signal ID"})
		}

		signal, err := s.GetSignal(c.Request().Context(), signalID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get signal"})
		}

		return c.JSON(http.StatusOK, signal)
	}
}

func GetSignalTracks(s *application.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		signalIDstr := c.Param("id")
		if signalIDstr == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Empty signal ID"})
		}

		signalID, err := strconv.Atoi(signalIDstr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid signal ID"})
		}

		tracks, err := s.GetSignalTracks(c.Request().Context(), signalID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to list signal tracks"})
		}

		return c.JSON(http.StatusOK, tracks)
	}
}

func ListSignalHandler(s *application.Service) echo.HandlerFunc {
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

		signals, nextPage, err := s.ListSignals(c.Request().Context(), limit, page)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to list signal"})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"signals":   signals,
			"next_page": nextPage,
		})
	}
}

func UpdateSignalHandler(s *application.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var signal domain.Signal
		if err := c.Bind(&signal); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		if err := s.UpdateSignal(c.Request().Context(), &signal); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update signal"})
		}

		return c.JSON(http.StatusOK, signal)
	}
}

func DeleteSignalHandler(s *application.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		signalIDstr := c.Param("id")
		if signalIDstr == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Empty signal ID"})
		}

		signalID, err := strconv.Atoi(signalIDstr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid signal ID"})
		}

		err = s.DeleteSignal(c.Request().Context(), signalID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete signal"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Deleted successfully"})
	}
}
