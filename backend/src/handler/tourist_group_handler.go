package handler

import (
	"errors"
	"net/http"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type TouristGroupHandler struct {
	service *service.TouristGroupService
}

func NewTouristGroupHandler(service *service.TouristGroupService) *TouristGroupHandler {
	return &TouristGroupHandler{
		service: service,
	}
}

func (h *TouristGroupHandler) Create(c echo.Context) error {
	var req dto.CreateTouristGroupRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Create(req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrTouristGroupInvalidName):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "name is required",
			})

		case errors.Is(err, service.ErrTouristGroupInvalidArrivalDate):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "arrival_date must be RFC3339 datetime",
			})

		case errors.Is(err, service.ErrTouristGroupInvalidDepartureDate):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "departure_date must be RFC3339 datetime",
			})

		case errors.Is(err, service.ErrTouristGroupInvalidDateRange):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "arrival_date must be before departure_date",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *TouristGroupHandler) GetByID(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	resp, err := h.service.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrTouristGroupNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "tourist group not found",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *TouristGroupHandler) List(c echo.Context) error {
	page := parseIntQuery(c, "page", 1)
	pageSize := parseIntQuery(c, "page_size", 20)

	resp, err := h.service.List(page, pageSize)
	if err != nil {
		return writeDatabaseError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *TouristGroupHandler) Update(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	var req dto.UpdateTouristGroupRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Update(id, req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrTouristGroupNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "tourist group not found",
			})

		case errors.Is(err, service.ErrTouristGroupInvalidName):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "name is required",
			})

		case errors.Is(err, service.ErrTouristGroupInvalidArrivalDate):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "arrival_date must be RFC3339 datetime",
			})

		case errors.Is(err, service.ErrTouristGroupInvalidDepartureDate):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "departure_date must be RFC3339 datetime",
			})

		case errors.Is(err, service.ErrTouristGroupInvalidDateRange):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "arrival_date must be before departure_date",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *TouristGroupHandler) Delete(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	err = h.service.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrTouristGroupNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "tourist group not found",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.NoContent(http.StatusNoContent)
}
