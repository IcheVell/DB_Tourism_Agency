package handler

import (
	"errors"
	"net/http"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type ExcursionBookingHandler struct {
	service *service.ExcursionBookingService
}

func NewExcursionBookingHandler(service *service.ExcursionBookingService) *ExcursionBookingHandler {
	return &ExcursionBookingHandler{
		service: service,
	}
}

func (h *ExcursionBookingHandler) Create(c echo.Context) error {
	var req dto.CreateExcursionBookingRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Create(req)
	if err != nil {
		return h.writeExcursionBookingError(c, err)
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *ExcursionBookingHandler) GetByID(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	resp, err := h.service.GetByID(id)
	if err != nil {
		return h.writeExcursionBookingError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *ExcursionBookingHandler) List(c echo.Context) error {
	page := parseIntQuery(c, "page", 1)
	pageSize := parseIntQuery(c, "page_size", 20)

	resp, err := h.service.List(page, pageSize)
	if err != nil {
		return writeDatabaseError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *ExcursionBookingHandler) Update(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	var req dto.UpdateExcursionBookingRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Update(id, req)
	if err != nil {
		return h.writeExcursionBookingError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *ExcursionBookingHandler) Delete(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	if err := h.service.Delete(id); err != nil {
		return h.writeExcursionBookingError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *ExcursionBookingHandler) writeExcursionBookingError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrExcursionBookingNotFound):
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "excursion booking not found",
		})

	case errors.Is(err, service.ErrExcursionBookingInvalidBookedAt):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "booked_at must be RFC3339 datetime",
		})

	case errors.Is(err, service.ErrExcursionBookingInvalidStatus):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "status must be booked, visited or cancelled",
		})

	case errors.Is(err, service.ErrExcursionBookingInvalidTouristRating):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "tourist_rating must be between 1 and 5",
		})

	case errors.Is(err, service.ErrExcursionBookingRatingRequiredForVisited):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "visited booking requires tourist_rating",
		})

	case errors.Is(err, service.ErrExcursionBookingRatingOnlyForVisited):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "tourist_rating is allowed only for visited status",
		})

	case errors.Is(err, service.ErrExcursionBookingInvalidScheduleID):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "excursion_schedule_id must be positive",
		})

	case errors.Is(err, service.ErrExcursionBookingInvalidGroupMemberID):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "group_member_id must be positive",
		})

	default:
		return writeDatabaseError(c, err)
	}
}
