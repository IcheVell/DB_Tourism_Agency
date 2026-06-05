package handler

import (
	"errors"
	"net/http"
	"strconv"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type PermissionHandler struct {
	permissionService *service.PermissionService
}

func NewPermissionHandler(permissionService *service.PermissionService) *PermissionHandler {
	return &PermissionHandler{
		permissionService: permissionService,
	}
}

func (h *PermissionHandler) Create(c echo.Context) error {
	var request dto.CreatePermissionRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body",
		})
	}

	permission, err := h.permissionService.Create(request)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusCreated, permission)
}

func (h *PermissionHandler) FindByID(c echo.Context) error {
	id, err := parsePermissionID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid permission id",
		})
	}

	permission, err := h.permissionService.FindByID(id)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, permission)
}

func (h *PermissionHandler) FindAll(c echo.Context) error {
	page := parsePermissionIntQuery(c, "page", 1)
	limit := parsePermissionIntQuery(c, "limit", 20)

	permissions, err := h.permissionService.FindAll(page, limit)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, permissions)
}

func (h *PermissionHandler) Update(c echo.Context) error {
	id, err := parsePermissionID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid permission id",
		})
	}

	var request dto.UpdatePermissionRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body",
		})
	}

	permission, err := h.permissionService.Update(id, request)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, permission)
}

func (h *PermissionHandler) Delete(c echo.Context) error {
	id, err := parsePermissionID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid permission id",
		})
	}

	if err := h.permissionService.Delete(id); err != nil {
		return h.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *PermissionHandler) handleError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrPermissionNotFound):
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Message: "permission not found",
		})

	case errors.Is(err, service.ErrPermissionAlreadyExists):
		return c.JSON(http.StatusConflict, dto.ErrorResponse{
			Message: "permission already exists",
		})

	case errors.Is(err, service.ErrPermissionHasRoles):
		return c.JSON(http.StatusConflict, dto.ErrorResponse{
			Message: "permission has roles",
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

func parsePermissionID(c echo.Context) (int64, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}

	if id <= 0 {
		return 0, strconv.ErrSyntax
	}

	return id, nil
}

func parsePermissionIntQuery(c echo.Context, name string, defaultValue int) int {
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
