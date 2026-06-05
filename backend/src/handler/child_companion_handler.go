package handler

import (
	"errors"
	"net/http"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type ChildCompanionHandler struct {
	service *service.ChildCompanionService
}

func NewChildCompanionHandler(service *service.ChildCompanionService) *ChildCompanionHandler {
	return &ChildCompanionHandler{
		service: service,
	}
}

func (h *ChildCompanionHandler) Create(c echo.Context) error {
	var req dto.CreateChildCompanionRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Create(req)
	if err != nil {
		return h.writeChildCompanionError(c, err)
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *ChildCompanionHandler) GetByIDs(c echo.Context) error {
	parentGroupMemberID, err := parseIDParam(c, "parent_group_member_id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid parent_group_member_id",
		})
	}

	childGroupMemberID, err := parseIDParam(c, "child_group_member_id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid child_group_member_id",
		})
	}

	resp, err := h.service.GetByIDs(parentGroupMemberID, childGroupMemberID)
	if err != nil {
		return h.writeChildCompanionError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *ChildCompanionHandler) List(c echo.Context) error {
	page := parseIntQuery(c, "page", 1)
	pageSize := parseIntQuery(c, "page_size", 20)

	resp, err := h.service.List(page, pageSize)
	if err != nil {
		return writeDatabaseError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *ChildCompanionHandler) Update(c echo.Context) error {
	parentGroupMemberID, err := parseIDParam(c, "parent_group_member_id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid parent_group_member_id",
		})
	}

	childGroupMemberID, err := parseIDParam(c, "child_group_member_id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid child_group_member_id",
		})
	}

	var req dto.UpdateChildCompanionRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Update(parentGroupMemberID, childGroupMemberID, req)
	if err != nil {
		return h.writeChildCompanionError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *ChildCompanionHandler) Delete(c echo.Context) error {
	parentGroupMemberID, err := parseIDParam(c, "parent_group_member_id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid parent_group_member_id",
		})
	}

	childGroupMemberID, err := parseIDParam(c, "child_group_member_id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid child_group_member_id",
		})
	}

	if err := h.service.Delete(parentGroupMemberID, childGroupMemberID); err != nil {
		return h.writeChildCompanionError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *ChildCompanionHandler) writeChildCompanionError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrChildCompanionNotFound):
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "child companion not found",
		})

	case errors.Is(err, service.ErrChildCompanionInvalidAdultID):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "parent_group_member_id must be positive",
		})

	case errors.Is(err, service.ErrChildCompanionInvalidChildID):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "child_group_member_id must be positive",
		})

	case errors.Is(err, service.ErrChildCompanionAdultEqualsChild):
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "parent_group_member_id and child_group_member_id must be different",
		})

	default:
		return writeDatabaseError(c, err)
	}
}
