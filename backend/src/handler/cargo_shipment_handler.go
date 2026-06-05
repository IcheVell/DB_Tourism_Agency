package handler

import (
	"errors"
	"net/http"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type CargoShipmentHandler struct {
	service *service.CargoShipmentService
}

func NewCargoShipmentHandler(service *service.CargoShipmentService) *CargoShipmentHandler {
	return &CargoShipmentHandler{
		service: service,
	}
}

func (h *CargoShipmentHandler) Create(c echo.Context) error {
	var req dto.CreateCargoShipmentRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Create(req)
	if err != nil {
		return h.writeCargoShipmentError(c, err)
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *CargoShipmentHandler) GetByID(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	resp, err := h.service.GetByID(id)
	if err != nil {
		return h.writeCargoShipmentError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *CargoShipmentHandler) List(c echo.Context) error {
	page := parseIntQuery(c, "page", 1)
	pageSize := parseIntQuery(c, "page_size", 20)

	resp, err := h.service.List(page, pageSize)
	if err != nil {
		return writeDatabaseError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *CargoShipmentHandler) Update(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	var req dto.UpdateCargoShipmentRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Update(id, req)
	if err != nil {
		return h.writeCargoShipmentError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *CargoShipmentHandler) Delete(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	if err := h.service.Delete(id); err != nil {
		return h.writeCargoShipmentError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *CargoShipmentHandler) writeCargoShipmentError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrCargoShipmentNotFound):
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "cargo shipment not found",
		})

	case errors.Is(err, service.ErrCargoShipmentInvalidStatus):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "status is required",
		})

	case errors.Is(err, service.ErrCargoShipmentInvalidShippedAt):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "shipped_at must be RFC3339 datetime",
		})

	case errors.Is(err, service.ErrCargoShipmentInvalidCargoStatementID):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "cargo_statement_id must be positive",
		})

	case errors.Is(err, service.ErrCargoShipmentInvalidFlightID):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "flight_id must be positive",
		})

	default:
		return writeDatabaseError(c, err)
	}
}
