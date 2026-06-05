package handler

import (
	"errors"
	"net/http"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type FlightHandler struct {
	service *service.FlightService
}

func NewFlightHandler(service *service.FlightService) *FlightHandler {
	return &FlightHandler{
		service: service,
	}
}

func (h *FlightHandler) Create(c echo.Context) error {
	var req dto.CreateFlightRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Create(req)
	if err != nil {
		return h.writeFlightError(c, err)
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *FlightHandler) GetByID(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	resp, err := h.service.GetByID(id)
	if err != nil {
		return h.writeFlightError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *FlightHandler) List(c echo.Context) error {
	page := parseIntQuery(c, "page", 1)
	pageSize := parseIntQuery(c, "page_size", 20)

	resp, err := h.service.List(page, pageSize)
	if err != nil {
		return writeDatabaseError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *FlightHandler) Update(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	var req dto.UpdateFlightRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Update(id, req)
	if err != nil {
		return h.writeFlightError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *FlightHandler) Delete(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	if err := h.service.Delete(id); err != nil {
		return h.writeFlightError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *FlightHandler) writeFlightError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrFlightNotFound):
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "flight not found",
		})

	case errors.Is(err, service.ErrFlightInvalidCapacity):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "capacity must be positive",
		})

	case errors.Is(err, service.ErrFlightInvalidFlightDate):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "flight_date must be YYYY-MM-DD date",
		})

	case errors.Is(err, service.ErrFlightInvalidFlightTypeID):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "flight_type_id must be positive",
		})

	default:
		return writeDatabaseError(c, err)
	}
}
