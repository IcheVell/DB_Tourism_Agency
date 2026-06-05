package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

type ReportHandler struct {
	reportService *service.ReportService
}

func NewReportHandler(reportService *service.ReportService) *ReportHandler {
	return &ReportHandler{
		reportService: reportService,
	}
}

func (h *ReportHandler) CustomsTourists(c echo.Context) error {
	categoryID := parseReportInt64Query(c, "category_id", 0)

	items, err := h.reportService.CustomsTourists(categoryID)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, items)
}

func (h *ReportHandler) AccommodationList(c echo.Context) error {
	hotelID := parseReportInt64Query(c, "hotel_id", 0)
	categoryID := parseReportInt64Query(c, "category_id", 0)

	items, err := h.reportService.AccommodationList(hotelID, categoryID)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, items)
}

func (h *ReportHandler) TouristCount(c echo.Context) error {
	from, err := parseReportOptionalTime(c.QueryParam("from"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid from"})
	}

	to, err := parseReportOptionalTime(c.QueryParam("to"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid to"})
	}

	categoryID := parseReportInt64Query(c, "category_id", 0)

	report, err := h.reportService.TouristCount(from, to, categoryID)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, report)
}

func (h *ReportHandler) TouristInfo(c echo.Context) error {
	touristID, err := parseReportInt64Param(c, "tourist_id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid tourist id"})
	}

	report, err := h.reportService.TouristInfo(touristID)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, report)
}

func (h *ReportHandler) HotelOccupancy(c echo.Context) error {
	from, err := parseReportOptionalTime(c.QueryParam("from"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid from"})
	}

	to, err := parseReportOptionalTime(c.QueryParam("to"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid to"})
	}

	items, err := h.reportService.HotelOccupancy(from, to)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, items)
}

func (h *ReportHandler) ExcursionTouristCount(c echo.Context) error {
	from, err := parseReportOptionalTime(c.QueryParam("from"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid from"})
	}

	to, err := parseReportOptionalTime(c.QueryParam("to"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid to"})
	}

	report, err := h.reportService.ExcursionTouristCount(from, to)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, report)
}

func (h *ReportHandler) ExcursionAnalytics(c echo.Context) error {
	from, err := parseReportOptionalTime(c.QueryParam("from"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid from"})
	}

	to, err := parseReportOptionalTime(c.QueryParam("to"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid to"})
	}

	report, err := h.reportService.ExcursionAnalytics(from, to)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, report)
}

func (h *ReportHandler) FlightLoad(c echo.Context) error {
	flightID := parseReportInt64Query(c, "flight_id", 0)

	report, err := h.reportService.FlightLoad(flightID)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, report)
}

func (h *ReportHandler) WarehouseTurnover(c echo.Context) error {
	from, err := parseReportOptionalTime(c.QueryParam("from"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid from"})
	}

	to, err := parseReportOptionalTime(c.QueryParam("to"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid to"})
	}

	report, err := h.reportService.WarehouseTurnover(from, to)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, report)
}

func (h *ReportHandler) GroupFinancialReport(c echo.Context) error {
	groupID := parseReportInt64Query(c, "group_id", 0)
	categoryID := parseReportInt64Query(c, "category_id", 0)

	items, err := h.reportService.GroupFinancialReport(groupID, categoryID)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, items)
}

func (h *ReportHandler) IncomeExpense(c echo.Context) error {
	from, err := parseReportOptionalTime(c.QueryParam("from"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid from"})
	}

	to, err := parseReportOptionalTime(c.QueryParam("to"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid to"})
	}

	items, err := h.reportService.IncomeExpense(from, to)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, items)
}

func (h *ReportHandler) CargoTypeShare(c echo.Context) error {
	from, err := parseReportOptionalTime(c.QueryParam("from"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid from"})
	}

	to, err := parseReportOptionalTime(c.QueryParam("to"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid to"})
	}

	items, err := h.reportService.CargoTypeShare(from, to)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, items)
}

func (h *ReportHandler) Profitability(c echo.Context) error {
	from, err := parseReportOptionalTime(c.QueryParam("from"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid from"})
	}

	to, err := parseReportOptionalTime(c.QueryParam("to"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid to"})
	}

	report, err := h.reportService.Profitability(from, to)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, report)
}

func (h *ReportHandler) TouristCategoryRatio(c echo.Context) error {
	from, err := parseReportOptionalTime(c.QueryParam("from"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid from"})
	}

	to, err := parseReportOptionalTime(c.QueryParam("to"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid to"})
	}

	restCategoryID := parseReportInt64Query(c, "rest_category_id", 0)
	shopCategoryID := parseReportInt64Query(c, "shop_category_id", 0)

	report, err := h.reportService.TouristCategoryRatio(from, to, restCategoryID, shopCategoryID)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, report)
}

func (h *ReportHandler) FlightTourists(c echo.Context) error {
	flightID := parseReportInt64Query(c, "flight_id", 0)

	items, err := h.reportService.FlightTourists(flightID)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.JSON(http.StatusOK, items)
}

func (h *ReportHandler) handleError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrInvalidInput):
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid input"})

	case errors.Is(err, service.ErrReportTouristNotFound):
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "tourist not found"})

	case errors.Is(err, service.ErrReportFlightNotFound):
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "flight not found"})

	default:
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "internal server error"})
	}
}

func parseReportInt64Param(c echo.Context, name string) (int64, error) {
	value, err := strconv.ParseInt(c.Param(name), 10, 64)
	if err != nil {
		return 0, err
	}

	if value <= 0 {
		return 0, strconv.ErrSyntax
	}

	return value, nil
}

func parseReportInt64Query(c echo.Context, name string, defaultValue int64) int64 {
	value := c.QueryParam(name)
	if value == "" {
		return defaultValue
	}

	parsedValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue
	}

	return parsedValue
}

func parseReportOptionalTime(value string) (*time.Time, error) {
	if value == "" {
		return nil, nil
	}

	parsedValue, err := time.Parse(time.RFC3339, value)
	if err == nil {
		return &parsedValue, nil
	}

	parsedValue, err = time.Parse("2006-01-02", value)
	if err != nil {
		return nil, err
	}

	return &parsedValue, nil
}
