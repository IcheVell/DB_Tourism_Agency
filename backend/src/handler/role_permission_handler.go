package handler

import (
	"errors"
	"net/http"
	"strconv"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type RolePermissionHandler struct {
	rolePermissionService *service.RolePermissionService
}

func NewRolePermissionHandler(rolePermissionService *service.RolePermissionService) *RolePermissionHandler {
	return &RolePermissionHandler{
		rolePermissionService: rolePermissionService,
	}
}

func (h *RolePermissionHandler) Create(c echo.Context) error {
	var request dto.CreateRolePermissionRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body",
		})
	}

	rolePermission, err := h.rolePermissionService.Create(request)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusCreated, rolePermission)
}

func (h *RolePermissionHandler) FindByIDs(c echo.Context) error {
	roleID, permissionID, err := parseRolePermissionIDs(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid role permission ids",
		})
	}

	rolePermission, err := h.rolePermissionService.FindByIDs(roleID, permissionID)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, rolePermission)
}

func (h *RolePermissionHandler) FindAll(c echo.Context) error {
	page := parseRolePermissionIntQuery(c, "page", 1)
	limit := parseRolePermissionIntQuery(c, "limit", 20)

	var roleID *int64

	roleIDParam := c.QueryParam("role_id")
	if roleIDParam != "" {
		parsedRoleID, err := strconv.ParseInt(roleIDParam, 10, 64)
		if err != nil || parsedRoleID <= 0 {
			return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Message: "invalid role id",
			})
		}

		roleID = &parsedRoleID
	}

	rolePermissions, err := h.rolePermissionService.FindAll(page, limit, roleID)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, rolePermissions)
}

func (h *RolePermissionHandler) Update(c echo.Context) error {
	roleID, permissionID, err := parseRolePermissionIDs(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid role permission ids",
		})
	}

	var request dto.UpdateRolePermissionRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body",
		})
	}

	rolePermission, err := h.rolePermissionService.Update(roleID, permissionID, request)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, rolePermission)
}

func (h *RolePermissionHandler) Delete(c echo.Context) error {
	roleID, permissionID, err := parseRolePermissionIDs(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid role permission ids",
		})
	}

	if err := h.rolePermissionService.Delete(roleID, permissionID); err != nil {
		return h.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *RolePermissionHandler) handleError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrRolePermissionNotFound):
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Message: "role permission not found",
		})

	case errors.Is(err, service.ErrRolePermissionAlreadyExists):
		return c.JSON(http.StatusConflict, dto.ErrorResponse{
			Message: "role permission already exists",
		})

	case errors.Is(err, service.ErrRoleNotFound):
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "role not found",
		})

	case errors.Is(err, service.ErrPermissionNotFound):
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "permission not found",
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

func parseRolePermissionIDs(c echo.Context) (int64, int64, error) {
	roleID, err := strconv.ParseInt(c.Param("role_id"), 10, 64)
	if err != nil {
		return 0, 0, err
	}

	if roleID <= 0 {
		return 0, 0, strconv.ErrSyntax
	}

	permissionID, err := strconv.ParseInt(c.Param("permission_id"), 10, 64)
	if err != nil {
		return 0, 0, err
	}

	if permissionID <= 0 {
		return 0, 0, strconv.ErrSyntax
	}

	return roleID, permissionID, nil
}

func parseRolePermissionIntQuery(c echo.Context, name string, defaultValue int) int {
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
