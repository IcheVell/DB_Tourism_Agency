package handler

import (
	"errors"
	"net/http"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type IdentityDocumentHandler struct {
	service *service.IdentityDocumentService
}

func NewIdentityDocumentHandler(service *service.IdentityDocumentService) *IdentityDocumentHandler {
	return &IdentityDocumentHandler{
		service: service,
	}
}

func (h *IdentityDocumentHandler) Create(c echo.Context) error {
	var req dto.CreateIdentityDocumentRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Create(req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrIdentityDocumentInvalidType):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "document_type must be PASSPORT, BIRTH_CERTIFICATE or INTERNATIONAL_PASSPORT",
			})

		case errors.Is(err, service.ErrIdentityDocumentInvalidSeries):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "document_series is required",
			})

		case errors.Is(err, service.ErrIdentityDocumentInvalidNumber):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "document_number is required",
			})

		case errors.Is(err, service.ErrIdentityDocumentInvalidIssuedBy):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "issued_by is required",
			})

		case errors.Is(err, service.ErrIdentityDocumentInvalidIssueDate):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "issue_date must be in YYYY-MM-DD format",
			})

		case errors.Is(err, service.ErrIdentityDocumentInvalidExpirationDate):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "expiration_date must be in YYYY-MM-DD format and greater than issue_date",
			})

		case errors.Is(err, service.ErrIdentityDocumentInvalidCitizenship):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "citizenship is required",
			})

		case errors.Is(err, service.ErrIdentityDocumentInvalidTouristID):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "tourist_id must be positive",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *IdentityDocumentHandler) GetByID(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	resp, err := h.service.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrIdentityDocumentNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "identity document not found",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *IdentityDocumentHandler) List(c echo.Context) error {
	page := parseIntQuery(c, "page", 1)
	pageSize := parseIntQuery(c, "page_size", 20)

	resp, err := h.service.List(page, pageSize)
	if err != nil {
		return writeDatabaseError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *IdentityDocumentHandler) Update(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	var req dto.UpdateIdentityDocumentRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Update(id, req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrIdentityDocumentNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "identity document not found",
			})

		case errors.Is(err, service.ErrIdentityDocumentInvalidType):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "document_type must be PASSPORT, BIRTH_CERTIFICATE or INTERNATIONAL_PASSPORT",
			})

		case errors.Is(err, service.ErrIdentityDocumentInvalidSeries):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "document_series is required",
			})

		case errors.Is(err, service.ErrIdentityDocumentInvalidNumber):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "document_number is required",
			})

		case errors.Is(err, service.ErrIdentityDocumentInvalidIssuedBy):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "issued_by is required",
			})

		case errors.Is(err, service.ErrIdentityDocumentInvalidIssueDate):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "issue_date must be in YYYY-MM-DD format",
			})

		case errors.Is(err, service.ErrIdentityDocumentInvalidExpirationDate):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "expiration_date must be in YYYY-MM-DD format and greater than issue_date",
			})

		case errors.Is(err, service.ErrIdentityDocumentInvalidCitizenship):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "citizenship is required",
			})

		case errors.Is(err, service.ErrIdentityDocumentInvalidTouristID):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "tourist_id must be positive",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *IdentityDocumentHandler) Delete(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	err = h.service.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrIdentityDocumentNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "identity document not found",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.NoContent(http.StatusNoContent)
}
