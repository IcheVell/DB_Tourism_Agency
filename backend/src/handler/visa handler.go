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
			"error": "Некорректное тело запроса",
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
			"error": "Некорректный id",
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
			"error": "Некорректный id",
		})
	}

	var req dto.UpdateVisaRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Некорректное тело запроса",
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
			"error": "Некорректный id",
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
			"error": "Виза не найдена",
		})

	case errors.Is(err, service.ErrVisaInvalidNumber):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Укажите номер визы",
		})

	case errors.Is(err, service.ErrVisaInvalidDestinationCountry):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Укажите страну назначения",
		})

	case errors.Is(err, service.ErrVisaInvalidStatus):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Статус визы должен быть draft, submitted, approved, rejected, issued, cancelled или expired",
		})

	case errors.Is(err, service.ErrVisaInvalidTouristID):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Выберите туриста",
		})

	case errors.Is(err, service.ErrVisaInvalidSubmittedAt):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Дата подачи должна быть корректной датой и временем",
		})

	case errors.Is(err, service.ErrVisaInvalidDecisionAt):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Дата решения должна быть корректной датой и временем",
		})

	case errors.Is(err, service.ErrVisaInvalidIssuedAt):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Дата выдачи визы должна быть корректной датой и временем",
		})

	case errors.Is(err, service.ErrVisaInvalidValidFrom):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Дата начала действия визы должна быть в формате YYYY-MM-DD",
		})

	case errors.Is(err, service.ErrVisaInvalidValidUntil):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Дата окончания действия визы должна быть в формате YYYY-MM-DD",
		})

	case errors.Is(err, service.ErrVisaInvalidTiming):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Даты визы указаны в неправильном порядке",
		})

	case errors.Is(err, service.ErrVisaInvalidValidityPeriod):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Дата окончания действия визы должна быть позже даты начала",
		})

	case errors.Is(err, service.ErrVisaIssuedFieldsRequired):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Для выданной визы нужны номер, дата выдачи, начало и окончание действия",
		})

	default:
		return writeDatabaseError(c, err)
	}
}
