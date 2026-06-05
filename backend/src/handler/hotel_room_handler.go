package handler

import (
	"errors"
	"net/http"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type HotelRoomHandler struct {
	service *service.HotelRoomService
}

func NewHotelRoomHandler(service *service.HotelRoomService) *HotelRoomHandler {
	return &HotelRoomHandler{
		service: service,
	}
}

func (h *HotelRoomHandler) Create(c echo.Context) error {
	var req dto.CreateHotelRoomRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Create(req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrHotelRoomInvalidNumber):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "room_number must be positive",
			})

		case errors.Is(err, service.ErrHotelRoomInvalidCapacity):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "capacity must be positive",
			})

		case errors.Is(err, service.ErrHotelRoomInvalidPrice):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "price must be positive",
			})

		case errors.Is(err, service.ErrHotelRoomInvalidHotelID):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "hotel_id must be positive",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *HotelRoomHandler) GetByID(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	resp, err := h.service.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrHotelRoomNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "hotel room not found",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HotelRoomHandler) List(c echo.Context) error {
	page := parseIntQuery(c, "page", 1)
	pageSize := parseIntQuery(c, "page_size", 20)

	resp, err := h.service.List(page, pageSize)
	if err != nil {
		return writeDatabaseError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HotelRoomHandler) Update(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	var req dto.UpdateHotelRoomRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.service.Update(id, req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrHotelRoomInvalidNumber):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "room_number must be positive",
			})

		case errors.Is(err, service.ErrHotelRoomInvalidCapacity):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "capacity must be positive",
			})

		case errors.Is(err, service.ErrHotelRoomInvalidPrice):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "price must be positive",
			})

		case errors.Is(err, service.ErrHotelRoomInvalidHotelID):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "hotel_id must be positive",
			})

		case errors.Is(err, service.ErrHotelRoomNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "hotel room not found",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HotelRoomHandler) Delete(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	err = h.service.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrHotelRoomNotFound):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "hotel room not found",
			})

		default:
			return writeDatabaseError(c, err)
		}
	}

	return c.NoContent(http.StatusNoContent)
}
