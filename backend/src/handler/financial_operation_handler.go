package handler

import (
	"errors"
	"net/http"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type FinancialOperationHandler struct {
	service *service.FinancialOperationService
}

func NewFinancialOperationHandler(service *service.FinancialOperationService) *FinancialOperationHandler {
	return &FinancialOperationHandler{
		service: service,
	}
}

func (h *FinancialOperationHandler) Create(c echo.Context) error {
	var req dto.CreateFinancialOperationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Create(req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrFinancialOperationInvalidAmount):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "amount must be positive",
			})

		case errors.Is(err, service.ErrFinancialOperationInvalidCategoryID):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "financial_category_id must be positive",
			})

		case errors.Is(err, service.ErrFinancialOperationInvalidSource):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "financial operation must have exactly one source",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *FinancialOperationHandler) GetByID(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	resp, err := h.service.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrFinancialOperationNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "financial operation not found",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *FinancialOperationHandler) List(c echo.Context) error {
	page := parseIntQuery(c, "page", 1)
	pageSize := parseIntQuery(c, "page_size", 20)

	resp, err := h.service.List(page, pageSize)
	if err != nil {
		return writeDatabaseError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *FinancialOperationHandler) Update(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	var req dto.UpdateFinancialOperationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Update(id, req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrFinancialOperationInvalidAmount):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "amount must be positive",
			})

		case errors.Is(err, service.ErrFinancialOperationInvalidCategoryID):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "financial_category_id must be positive",
			})

		case errors.Is(err, service.ErrFinancialOperationInvalidSource):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "financial operation must have exactly one source",
			})

		case errors.Is(err, service.ErrFinancialOperationNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "financial operation not found",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *FinancialOperationHandler) Delete(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	err = h.service.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrFinancialOperationNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "financial operation not found",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.NoContent(http.StatusNoContent)
}
