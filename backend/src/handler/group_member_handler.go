package handler

import (
	"errors"
	"net/http"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type GroupMemberHandler struct {
	service *service.GroupMemberService
}

func NewGroupMemberHandler(service *service.GroupMemberService) *GroupMemberHandler {
	return &GroupMemberHandler{
		service: service,
	}
}

func (h *GroupMemberHandler) Create(c echo.Context) error {
	var req dto.CreateGroupMemberRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Create(req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrGroupMemberInvalidTouristGroupID):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "tourist_group_id must be positive",
			})

		case errors.Is(err, service.ErrGroupMemberInvalidTouristCategoryID):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "tourist_category_id must be positive",
			})

		case errors.Is(err, service.ErrGroupMemberInvalidTouristID):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "tourist_id must be positive",
			})

		case errors.Is(err, service.ErrGroupMemberInvalidDesiredHotelID):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "desired_hotel_id must be positive",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *GroupMemberHandler) GetByID(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	resp, err := h.service.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrGroupMemberNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "group member not found",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *GroupMemberHandler) List(c echo.Context) error {
	page := parseIntQuery(c, "page", 1)
	pageSize := parseIntQuery(c, "page_size", 20)

	resp, err := h.service.List(page, pageSize)
	if err != nil {
		return writeDatabaseError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *GroupMemberHandler) Update(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	var req dto.UpdateGroupMemberRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Update(id, req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrGroupMemberNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "group member not found",
			})

		case errors.Is(err, service.ErrGroupMemberInvalidTouristGroupID):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "tourist_group_id must be positive",
			})

		case errors.Is(err, service.ErrGroupMemberInvalidTouristCategoryID):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "tourist_category_id must be positive",
			})

		case errors.Is(err, service.ErrGroupMemberInvalidTouristID):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "tourist_id must be positive",
			})

		case errors.Is(err, service.ErrGroupMemberInvalidDesiredHotelID):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "desired_hotel_id must be positive",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *GroupMemberHandler) Delete(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	err = h.service.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrGroupMemberNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "group member not found",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.NoContent(http.StatusNoContent)
}
