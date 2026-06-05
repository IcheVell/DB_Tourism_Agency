package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterRoutes(group *echo.Group) {
	users := group.Group("/users")

	users.POST("", h.Create)
	users.GET("", h.List)
	users.GET("/:id", h.GetByID)
	users.PATCH("/:id", h.Update)
	users.DELETE("/:id", h.Delete)
}

func (h *UserHandler) Create(c echo.Context) error {
	var req dto.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body",
		})
	}

	user, err := h.userService.Create(c.Request().Context(), req)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetByID(c echo.Context) error {
	id, err := parseUint64Param(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid user id",
		})
	}

	user, err := h.userService.GetByID(c.Request().Context(), id)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) List(c echo.Context) error {
	page := parseIntQuery(c, "page", 1)
	limit := parseIntQuery(c, "limit", 20)

	users, err := h.userService.List(c.Request().Context(), page, limit)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Update(c echo.Context) error {
	id, err := parseUint64Param(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid user id",
		})
	}

	var req dto.UpdateUserRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body",
		})
	}

	user, err := h.userService.Update(c.Request().Context(), id, req)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Delete(c echo.Context) error {
	id, err := parseUint64Param(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid user id",
		})
	}

	if err := h.userService.Delete(c.Request().Context(), id); err != nil {
		return h.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *UserHandler) handleError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrUserNotFound):
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Message: "user not found",
		})

	case errors.Is(err, service.ErrRoleNotFound):
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "role not found",
		})

	case errors.Is(err, service.ErrLoginAlreadyUsed):
		return c.JSON(http.StatusConflict, dto.ErrorResponse{
			Message: "login already used",
		})

	case errors.Is(err, service.ErrEmailAlreadyUsed):
		return c.JSON(http.StatusConflict, dto.ErrorResponse{
			Message: "email already used",
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

func parseUint64Param(c echo.Context, name string) (uint64, error) {
	value := c.Param(name)

	id, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, strconv.ErrSyntax
	}

	return id, nil
}
