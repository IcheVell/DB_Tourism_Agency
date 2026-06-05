package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type MeHandler struct {
	meService *service.MeService
}

func NewMeHandler(meService *service.MeService) *MeHandler {
	return &MeHandler{
		meService: meService,
	}
}

func (h *MeHandler) Tours(c echo.Context) error {
	userID, err := getCurrentUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "unauthorized",
		})
	}

	page := parseMeIntQuery(c, "page", 1)
	limit := parseMeIntQuery(c, "limit", 20)

	result, err := h.meService.Tours(userID, page, limit)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *MeHandler) Visas(c echo.Context) error {
	userID, err := getCurrentUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "unauthorized",
		})
	}

	page := parseMeIntQuery(c, "page", 1)
	limit := parseMeIntQuery(c, "limit", 20)

	result, err := h.meService.Visas(userID, page, limit)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *MeHandler) Accommodations(c echo.Context) error {
	userID, err := getCurrentUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "unauthorized",
		})
	}

	page := parseMeIntQuery(c, "page", 1)
	limit := parseMeIntQuery(c, "limit", 20)

	result, err := h.meService.Accommodations(userID, page, limit)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *MeHandler) Excursions(c echo.Context) error {
	userID, err := getCurrentUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "unauthorized",
		})
	}

	page := parseMeIntQuery(c, "page", 1)
	limit := parseMeIntQuery(c, "limit", 20)

	result, err := h.meService.Excursions(userID, page, limit)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *MeHandler) Cargo(c echo.Context) error {
	userID, err := getCurrentUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "unauthorized",
		})
	}

	page := parseMeIntQuery(c, "page", 1)
	limit := parseMeIntQuery(c, "limit", 20)

	result, err := h.meService.Cargo(userID, page, limit)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *MeHandler) handleError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrMeTouristNotLinked):
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Message: "tourist profile is not linked to user",
		})

	case errors.Is(err, service.ErrInvalidInput):
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid input",
		})

	default:
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "internal server error",
		})
	}
}

func getCurrentUserID(c echo.Context) (int64, error) {
	keys := []string{
		"user_id",
		"userID",
		"id",
	}

	for _, key := range keys {
		value := c.Get(key)
		if value == nil {
			continue
		}

		userID, err := parseMeAnyInt64(value)
		if err == nil && userID > 0 {
			return userID, nil
		}
	}

	if claims, ok := c.Get("claims").(map[string]any); ok {
		for _, key := range keys {
			value, exists := claims[key]
			if !exists {
				continue
			}

			userID, err := parseMeAnyInt64(value)
			if err == nil && userID > 0 {
				return userID, nil
			}
		}
	}

	return 0, errors.New("user id not found in context")
}

func parseMeAnyInt64(value any) (int64, error) {
	switch typedValue := value.(type) {
	case int:
		return int64(typedValue), nil
	case int8:
		return int64(typedValue), nil
	case int16:
		return int64(typedValue), nil
	case int32:
		return int64(typedValue), nil
	case int64:
		return typedValue, nil
	case uint:
		return int64(typedValue), nil
	case uint8:
		return int64(typedValue), nil
	case uint16:
		return int64(typedValue), nil
	case uint32:
		return int64(typedValue), nil
	case uint64:
		if typedValue > uint64(^uint(0)>>1) {
			return 0, fmt.Errorf("uint64 overflow")
		}

		return int64(typedValue), nil
	case float64:
		return int64(typedValue), nil
	case string:
		return strconv.ParseInt(typedValue, 10, 64)
	default:
		return 0, fmt.Errorf("unsupported user id type")
	}
}

func parseMeIntQuery(c echo.Context, name string, defaultValue int) int {
	value := c.QueryParam(name)
	if value == "" {
		return defaultValue
	}

	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return parsedValue
}

func (h *MeHandler) IdentityDocument(c echo.Context) error {
	userID, err := getCurrentUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "unauthorized",
		})
	}

	document, err := h.meService.IdentityDocument(userID)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, document)
}

func (h *MeHandler) CreateIdentityDocument(c echo.Context) error {
	userID, err := getCurrentUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "unauthorized",
		})
	}

	var request dto.CreateMeIdentityDocumentRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body",
		})
	}

	document, err := h.meService.CreateIdentityDocument(userID, request)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusCreated, document)
}

func (h *MeHandler) UpdateIdentityDocument(c echo.Context) error {
	userID, err := getCurrentUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "unauthorized",
		})
	}

	var request dto.UpdateMeIdentityDocumentRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body",
		})
	}

	document, err := h.meService.UpdateIdentityDocument(userID, request)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, document)
}

func (h *MeHandler) CreateExcursionBooking(c echo.Context) error {
	userID, err := getCurrentUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "unauthorized",
		})
	}

	var request dto.CreateMeExcursionBookingRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body",
		})
	}

	booking, err := h.meService.CreateExcursionBooking(userID, request)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusCreated, booking)
}
