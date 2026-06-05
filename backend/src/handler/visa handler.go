package handler

import (
	"errors"
	"net/http"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type VisaHandler struct {
	service *service.VisaService
}

func NewVisaHandler(service *service.VisaService) *VisaHandler {
	return &VisaHandler{
		service: service,
	}
}

func (h *VisaHandler) Create(c echo.Context) error {
	var req dto.CreateVisaRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Create(req)
	if err != nil {
		return h.writeVisaError(c, err)
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *VisaHandler) GetByID(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	resp, err := h.service.GetByID(id)
	if err != nil {
		return h.writeVisaError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *VisaHandler) List(c echo.Context) error {
	page := parseIntQuery(c, "page", 1)
	pageSize := parseIntQuery(c, "page_size", 20)

	resp, err := h.service.List(page, pageSize)
	if err != nil {
		return writeDatabaseError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *VisaHandler) Update(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	var req dto.UpdateVisaRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Update(id, req)
	if err != nil {
		return h.writeVisaError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *VisaHandler) Delete(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	if err := h.service.Delete(id); err != nil {
		return h.writeVisaError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *VisaHandler) writeVisaError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrVisaNotFound):
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "visa not found",
		})

	case errors.Is(err, service.ErrVisaInvalidNumber):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "number is required",
		})

	case errors.Is(err, service.ErrVisaInvalidDestinationCountry):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "destination_country is required",
		})

	case errors.Is(err, service.ErrVisaInvalidStatus):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "status must be draft, submitted, rejected, issued, cancelled or expired",
		})

	case errors.Is(err, service.ErrVisaInvalidTouristID):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "tourist_id must be positive",
		})

	case errors.Is(err, service.ErrVisaInvalidSubmittedAt):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "submitted_at must be RFC3339 datetime",
		})

	case errors.Is(err, service.ErrVisaInvalidDecisionAt):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "decision_at must be RFC3339 datetime",
		})

	case errors.Is(err, service.ErrVisaInvalidIssuedAt):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "issued_at must be RFC3339 datetime",
		})

	case errors.Is(err, service.ErrVisaInvalidValidFrom):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "valid_from must be YYYY-MM-DD date",
		})

	case errors.Is(err, service.ErrVisaInvalidValidUntil):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "valid_until must be YYYY-MM-DD date",
		})

	case errors.Is(err, service.ErrVisaInvalidTiming):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "visa timestamps are inconsistent",
		})

	case errors.Is(err, service.ErrVisaInvalidValidityPeriod):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "valid_from must be before valid_until",
		})

	case errors.Is(err, service.ErrVisaIssuedFieldsRequired):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "issued visa requires number, issued_at, valid_from and valid_until",
		})

	default:
		return writeDatabaseError(c, err)
	}
}
