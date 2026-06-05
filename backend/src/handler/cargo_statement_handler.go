package handler

import (
	"errors"
	"net/http"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type CargoStatementHandler struct {
	service *service.CargoStatementService
}

func NewCargoStatementHandler(service *service.CargoStatementService) *CargoStatementHandler {
	return &CargoStatementHandler{
		service: service,
	}
}

func (h *CargoStatementHandler) Create(c echo.Context) error {
	var req dto.CreateCargoStatementRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Create(req)
	if err != nil {
		return h.writeCargoStatementError(c, err)
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *CargoStatementHandler) GetByID(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	resp, err := h.service.GetByID(id)
	if err != nil {
		return h.writeCargoStatementError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *CargoStatementHandler) List(c echo.Context) error {
	page := parseIntQuery(c, "page", 1)
	pageSize := parseIntQuery(c, "page_size", 20)

	resp, err := h.service.List(page, pageSize)
	if err != nil {
		return writeDatabaseError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *CargoStatementHandler) Update(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	var req dto.UpdateCargoStatementRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Update(id, req)
	if err != nil {
		return h.writeCargoStatementError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *CargoStatementHandler) Delete(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	if err := h.service.Delete(id); err != nil {
		return h.writeCargoStatementError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *CargoStatementHandler) writeCargoStatementError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrCargoStatementNotFound):
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "cargo statement not found",
		})

	case errors.Is(err, service.ErrCargoStatementInvalidStatus):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "status must be draft, weighed, packed, ready_for_shipment, shipped or cancelled",
		})

	case errors.Is(err, service.ErrCargoStatementInvalidGroupMemberID):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "group_member_id must be positive",
		})

	default:
		return writeDatabaseError(c, err)
	}
}
