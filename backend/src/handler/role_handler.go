package handler

import (
	"errors"
	"net/http"
	"strconv"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type RoleHandler struct {
	roleService *service.RoleService
}

func NewRoleHandler(roleService *service.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService: roleService,
	}
}

func (h *RoleHandler) Create(c echo.Context) error {
	var request dto.CreateRoleRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body",
		})
	}

	role, err := h.roleService.Create(request)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusCreated, role)
}

func (h *RoleHandler) FindByID(c echo.Context) error {
	id, err := parseRoleID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid role id",
		})
	}

	role, err := h.roleService.FindByID(id)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, role)
}

func (h *RoleHandler) FindAll(c echo.Context) error {
	page := parseRoleIntQuery(c, "page", 1)
	limit := parseRoleIntQuery(c, "limit", 20)

	roles, err := h.roleService.FindAll(page, limit)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, roles)
}

func (h *RoleHandler) Update(c echo.Context) error {
	id, err := parseRoleID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid role id",
		})
	}

	var request dto.UpdateRoleRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body",
		})
	}

	role, err := h.roleService.Update(id, request)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, role)
}

func (h *RoleHandler) Delete(c echo.Context) error {
	id, err := parseRoleID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid role id",
		})
	}

	if err := h.roleService.Delete(id); err != nil {
		return h.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *RoleHandler) handleError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrRoleNotFound):
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Message: "role not found",
		})

	case errors.Is(err, service.ErrRoleAlreadyExists):
		return c.JSON(http.StatusConflict, dto.ErrorResponse{
			Message: "role already exists",
		})

	case errors.Is(err, service.ErrRoleHasUsers):
		return c.JSON(http.StatusConflict, dto.ErrorResponse{
			Message: "role has users",
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

func parseRoleID(c echo.Context) (int64, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}

	if id <= 0 {
		return 0, strconv.ErrSyntax
	}

	return id, nil
}

func parseRoleIntQuery(c echo.Context, name string, defaultValue int) int {
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
