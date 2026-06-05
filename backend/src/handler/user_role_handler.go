package handler

import (
	"errors"
	"net/http"
	"strconv"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type UserRoleHandler struct {
	userRoleService *service.UserRoleService
}

func NewUserRoleHandler(userRoleService *service.UserRoleService) *UserRoleHandler {
	return &UserRoleHandler{
		userRoleService: userRoleService,
	}
}

func (h *UserRoleHandler) Create(c echo.Context) error {
	var request dto.CreateUserRoleRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body",
		})
	}

	userRole, err := h.userRoleService.Create(request)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusCreated, userRole)
}

func (h *UserRoleHandler) FindByIDs(c echo.Context) error {
	userID, roleID, err := parseUserRoleIDs(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid user role ids",
		})
	}

	userRole, err := h.userRoleService.FindByIDs(userID, roleID)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, userRole)
}

func (h *UserRoleHandler) FindByUserID(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil || userID <= 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid user id",
		})
	}

	userRole, err := h.userRoleService.FindByUserID(userID)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, userRole)
}

func (h *UserRoleHandler) FindAll(c echo.Context) error {
	page := parseUserRoleIntQuery(c, "page", 1)
	limit := parseUserRoleIntQuery(c, "limit", 20)

	var userID *int64

	userIDParam := c.QueryParam("user_id")
	if userIDParam != "" {
		parsedUserID, err := strconv.ParseInt(userIDParam, 10, 64)
		if err != nil || parsedUserID <= 0 {
			return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Message: "invalid user id",
			})
		}

		userID = &parsedUserID
	}

	userRoles, err := h.userRoleService.FindAll(page, limit, userID)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, userRoles)
}

func (h *UserRoleHandler) Update(c echo.Context) error {
	userID, roleID, err := parseUserRoleIDs(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid user role ids",
		})
	}

	var request dto.UpdateUserRoleRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body",
		})
	}

	userRole, err := h.userRoleService.Update(userID, roleID, request)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, userRole)
}

func (h *UserRoleHandler) Delete(c echo.Context) error {
	userID, roleID, err := parseUserRoleIDs(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid user role ids",
		})
	}

	if err := h.userRoleService.Delete(userID, roleID); err != nil {
		return h.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *UserRoleHandler) handleError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrUserRoleNotFound):
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Message: "user role not found",
		})

	case errors.Is(err, service.ErrUserRoleAlreadyExists):
		return c.JSON(http.StatusConflict, dto.ErrorResponse{
			Message: "user already has role",
		})

	case errors.Is(err, service.ErrUserNotFound):
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "user not found",
		})

	case errors.Is(err, service.ErrRoleNotFound):
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "role not found",
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

func parseUserRoleIDs(c echo.Context) (int64, int64, error) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return 0, 0, err
	}

	if userID <= 0 {
		return 0, 0, strconv.ErrSyntax
	}

	roleID, err := strconv.ParseInt(c.Param("role_id"), 10, 64)
	if err != nil {
		return 0, 0, err
	}

	if roleID <= 0 {
		return 0, 0, strconv.ErrSyntax
	}

	return userID, roleID, nil
}

func parseUserRoleIntQuery(c echo.Context, name string, defaultValue int) int {
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
