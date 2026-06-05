package handler

import (
	"errors"
	"net/http"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type ExcursionScheduleHandler struct {
	service *service.ExcursionScheduleService
}

func NewExcursionScheduleHandler(service *service.ExcursionScheduleService) *ExcursionScheduleHandler {
	return &ExcursionScheduleHandler{
		service: service,
	}
}

func (h *ExcursionScheduleHandler) Create(c echo.Context) error {
	var req dto.CreateExcursionScheduleRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Create(req)
	if err != nil {
		return h.writeExcursionScheduleError(c, err)
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *ExcursionScheduleHandler) GetByID(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	resp, err := h.service.GetByID(id)
	if err != nil {
		return h.writeExcursionScheduleError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *ExcursionScheduleHandler) List(c echo.Context) error {
	page := parseIntQuery(c, "page", 1)
	pageSize := parseIntQuery(c, "page_size", 20)

	resp, err := h.service.List(page, pageSize)
	if err != nil {
		return writeDatabaseError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *ExcursionScheduleHandler) Update(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	var req dto.UpdateExcursionScheduleRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Update(id, req)
	if err != nil {
		return h.writeExcursionScheduleError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *ExcursionScheduleHandler) Delete(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	if err := h.service.Delete(id); err != nil {
		return h.writeExcursionScheduleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *ExcursionScheduleHandler) writeExcursionScheduleError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrExcursionScheduleNotFound):
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "excursion schedule not found",
		})

	case errors.Is(err, service.ErrExcursionScheduleInvalidPrice):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "price must be positive",
		})

	case errors.Is(err, service.ErrExcursionScheduleInvalidStartTime):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "start_time must be RFC3339 datetime",
		})

	case errors.Is(err, service.ErrExcursionScheduleInvalidEndTime):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "end_time must be RFC3339 datetime",
		})

	case errors.Is(err, service.ErrExcursionScheduleInvalidTimeRange):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "start_time must be before end_time",
		})

	case errors.Is(err, service.ErrExcursionScheduleInvalidCapacity):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "capacity must be positive",
		})

	case errors.Is(err, service.ErrExcursionScheduleInvalidStatus):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "status must be planned, completed or cancelled",
		})

	case errors.Is(err, service.ErrExcursionScheduleInvalidAgencyID):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "excursion_agency_id must be positive",
		})

	case errors.Is(err, service.ErrExcursionScheduleInvalidExcursionID):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "excursion_id must be positive",
		})

	default:
		return writeDatabaseError(c, err)
	}
}
