package handler

import (
	"errors"
	"net/http"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type ExcursionAgencyHandler struct {
	service *service.ExcursionAgencyService
}

func NewExcursionAgencyHandler(service *service.ExcursionAgencyService) *ExcursionAgencyHandler {
	return &ExcursionAgencyHandler{
		service: service,
	}
}

func (h *ExcursionAgencyHandler) Create(c echo.Context) error {
	var req dto.CreateExcursionAgencyRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Create(req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrExcursionAgencyInvalidName):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "name is required",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *ExcursionAgencyHandler) GetByID(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	resp, err := h.service.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrExcursionAgencyNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "excursion agency not found",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *ExcursionAgencyHandler) List(c echo.Context) error {
	page := parseIntQuery(c, "page", 1)
	pageSize := parseIntQuery(c, "page_size", 20)

	resp, err := h.service.List(page, pageSize)
	if err != nil {
		return writeDatabaseError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *ExcursionAgencyHandler) Update(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	var req dto.UpdateExcursionAgencyRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Update(id, req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrExcursionAgencyInvalidName):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "name is required",
			})

		case errors.Is(err, service.ErrExcursionAgencyNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "excursion agency not found",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *ExcursionAgencyHandler) Delete(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	err = h.service.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrExcursionAgencyNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "excursion agency not found",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.NoContent(http.StatusNoContent)
}
