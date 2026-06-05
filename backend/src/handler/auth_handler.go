package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var request dto.RegisterRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body",
		})
	}

	response, err := h.authService.Register(request)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *AuthHandler) Login(c echo.Context) error {
	var request dto.LoginRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body",
		})
	}

	response, err := h.authService.Login(request)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) Me(c echo.Context) error {
	userID, err := getAuthCurrentUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "unauthorized",
		})
	}

	response, err := h.authService.Me(userID)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) handleError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrAuthInvalidInput):
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid input",
		})

	case errors.Is(err, service.ErrAuthInvalidCredentials):
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "invalid login or password",
		})

	case errors.Is(err, service.ErrAuthUserNotFound):
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Message: "user not found",
		})

	case errors.Is(err, service.ErrAuthLoginAlreadyUsed):
		return c.JSON(http.StatusConflict, dto.ErrorResponse{
			Message: "login already used",
		})

	case errors.Is(err, service.ErrAuthEmailAlreadyUsed):
		return c.JSON(http.StatusConflict, dto.ErrorResponse{
			Message: "email already used",
		})

	case errors.Is(err, service.ErrAuthTouristRoleNotFound):
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "tourist role not found",
		})

	default:
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "internal server error",
		})
	}
}

func getAuthCurrentUserID(c echo.Context) (int64, error) {
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

		userID, err := parseAuthAnyInt64(value)
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

			userID, err := parseAuthAnyInt64(value)
			if err == nil && userID > 0 {
				return userID, nil
			}
		}
	}

	if token, ok := c.Get("user").(*jwt.Token); ok {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			for _, key := range keys {
				value, exists := claims[key]
				if !exists {
					continue
				}

				userID, err := parseAuthAnyInt64(value)
				if err == nil && userID > 0 {
					return userID, nil
				}
			}
		}
	}

	return 0, errors.New("user id not found in context")
}

func parseAuthAnyInt64(value any) (int64, error) {
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
