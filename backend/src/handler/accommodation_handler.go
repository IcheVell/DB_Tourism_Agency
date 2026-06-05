package handler

import (
	"errors"
	"net/http"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type AccommodationHandler struct {
	service *service.AccommodationService
}

func NewAccommodationHandler(service *service.AccommodationService) *AccommodationHandler {
	return &AccommodationHandler{
		service: service,
	}
}

func (h *AccommodationHandler) Create(c echo.Context) error {
	var req dto.CreateAccommodationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Create(req)
	if err != nil {
		return h.writeAccommodationError(c, err)
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *AccommodationHandler) GetByID(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	resp, err := h.service.GetByID(id)
	if err != nil {
		return h.writeAccommodationError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *AccommodationHandler) List(c echo.Context) error {
	page := parseIntQuery(c, "page", 1)
	pageSize := parseIntQuery(c, "page_size", 20)

	resp, err := h.service.List(page, pageSize)
	if err != nil {
		return writeDatabaseError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *AccommodationHandler) Update(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	var req dto.UpdateAccommodationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Update(id, req)
	if err != nil {
		return h.writeAccommodationError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *AccommodationHandler) Delete(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	if err := h.service.Delete(id); err != nil {
		return h.writeAccommodationError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *AccommodationHandler) writeAccommodationError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrAccommodationNotFound):
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "accommodation not found",
		})

	case errors.Is(err, service.ErrAccommodationInvalidStatus):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "status must be reserved, checked_in, checked_out or cancelled",
		})

	case errors.Is(err, service.ErrAccommodationInvalidCheckInAt):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "check_in_at must be RFC3339 datetime",
		})

	case errors.Is(err, service.ErrAccommodationInvalidCheckOutAt):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "check_out_at must be RFC3339 datetime",
		})

	case errors.Is(err, service.ErrAccommodationInvalidDateRange):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "check_out_at must be after check_in_at",
		})

	case errors.Is(err, service.ErrAccommodationCheckOutRequired):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "checked_out accommodation requires check_out_at",
		})

	case errors.Is(err, service.ErrAccommodationInvalidGroupMemberID):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "group_member_id must be positive",
		})

	case errors.Is(err, service.ErrAccommodationInvalidHotelRoomID):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "hotel_room_id must be positive",
		})

	default:
		return writeDatabaseError(c, err)
	}
}
