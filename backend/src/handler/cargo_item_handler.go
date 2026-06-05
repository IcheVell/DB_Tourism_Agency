package handler

import (
	"errors"
	"net/http"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type CargoItemHandler struct {
	service *service.CargoItemService
}

func NewCargoItemHandler(service *service.CargoItemService) *CargoItemHandler {
	return &CargoItemHandler{
		service: service,
	}
}

func (h *CargoItemHandler) Create(c echo.Context) error {
	var req dto.CreateCargoItemRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Create(req)
	if err != nil {
		return h.writeCargoItemError(c, err)
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *CargoItemHandler) GetByID(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	resp, err := h.service.GetByID(id)
	if err != nil {
		return h.writeCargoItemError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *CargoItemHandler) List(c echo.Context) error {
	page := parseIntQuery(c, "page", 1)
	pageSize := parseIntQuery(c, "page_size", 20)

	resp, err := h.service.List(page, pageSize)
	if err != nil {
		return writeDatabaseError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *CargoItemHandler) Update(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	var req dto.UpdateCargoItemRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Update(id, req)
	if err != nil {
		return h.writeCargoItemError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *CargoItemHandler) Delete(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	if err := h.service.Delete(id); err != nil {
		return h.writeCargoItemError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *CargoItemHandler) writeCargoItemError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrCargoItemNotFound):
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "cargo item not found",
		})

	case errors.Is(err, service.ErrCargoItemInvalidItemNumber):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "item_number is required",
		})

	case errors.Is(err, service.ErrCargoItemInvalidWeightKg):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "weight_kg must be positive",
		})

	case errors.Is(err, service.ErrCargoItemInvalidVolumetricWeightKg):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "volumetric_weight_kg must be positive",
		})

	case errors.Is(err, service.ErrCargoItemInvalidPlacesCount):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "places_count must be positive",
		})

	case errors.Is(err, service.ErrCargoItemInvalidPackagedAt):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "packaged_at must be RFC3339 datetime",
		})

	case errors.Is(err, service.ErrCargoItemInvalidCargoTypeID):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "cargo_type_id must be positive",
		})

	case errors.Is(err, service.ErrCargoItemInvalidCargoStatementID):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "cargo_statement_id must be positive",
		})

	default:
		return writeDatabaseError(c, err)
	}
}
