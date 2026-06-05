package service

import (
	"errors"
	"time"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/repository"
)

var (
	ErrReportTouristNotFound = errors.New("report tourist not found")
	ErrReportFlightNotFound  = errors.New("report flight not found")
)

type ReportService struct {
	reportRepository *repository.ReportRepository
}

func NewReportService(reportRepository *repository.ReportRepository) *ReportService {
	return &ReportService{
		reportRepository: reportRepository,
	}
}

func (s *ReportService) CustomsTourists(categoryID int64) ([]dto.CustomsTouristReportItem, error) {
	return s.reportRepository.FindCustomsTourists(categoryID)
}

func (s *ReportService) AccommodationList(hotelID int64, categoryID int64) ([]dto.AccommodationReportItem, error) {
	return s.reportRepository.FindAccommodationList(hotelID, categoryID)
}

func (s *ReportService) TouristCount(from *time.Time, to *time.Time, categoryID int64) (*dto.TouristCountReportResponse, error) {
	if invalidPeriod(from, to) {
		return nil, ErrInvalidInput
	}

	count, err := s.reportRepository.CountTourists(from, to, categoryID)
	if err != nil {
		return nil, err
	}

	return &dto.TouristCountReportResponse{
		Count: count,
	}, nil
}

func (s *ReportService) TouristInfo(touristID int64) (*dto.TouristInfoReportResponse, error) {
	if touristID <= 0 {
		return nil, ErrInvalidInput
	}

	report, err := s.reportRepository.FindTouristInfo(touristID)
	if err != nil {
		return nil, err
	}

	if report == nil {
		return nil, ErrReportTouristNotFound
	}

	return report, nil
}

func (s *ReportService) HotelOccupancy(from *time.Time, to *time.Time) ([]dto.HotelOccupancyReportItem, error) {
	if invalidPeriod(from, to) {
		return nil, ErrInvalidInput
	}

	return s.reportRepository.FindHotelOccupancy(from, to)
}

func (s *ReportService) ExcursionTouristCount(from *time.Time, to *time.Time) (*dto.ExcursionTouristCountReportResponse, error) {
	if invalidPeriod(from, to) {
		return nil, ErrInvalidInput
	}

	count, err := s.reportRepository.CountExcursionTourists(from, to)
	if err != nil {
		return nil, err
	}

	return &dto.ExcursionTouristCountReportResponse{
		Count: count,
	}, nil
}

func (s *ReportService) ExcursionAnalytics(from *time.Time, to *time.Time) (*dto.ExcursionAnalyticsReportResponse, error) {
	if invalidPeriod(from, to) {
		return nil, ErrInvalidInput
	}

	return s.reportRepository.FindExcursionAnalytics(from, to)
}

func (s *ReportService) FlightLoad(flightID int64) (*dto.FlightLoadReportResponse, error) {
	if flightID <= 0 {
		return nil, ErrInvalidInput
	}

	report, err := s.reportRepository.FindFlightLoad(flightID)
	if err != nil {
		return nil, err
	}

	if report == nil {
		return nil, ErrReportFlightNotFound
	}

	return report, nil
}

func (s *ReportService) WarehouseTurnover(from *time.Time, to *time.Time) (*dto.WarehouseTurnoverReportResponse, error) {
	if invalidPeriod(from, to) {
		return nil, ErrInvalidInput
	}

	return s.reportRepository.FindWarehouseTurnover(from, to)
}

func (s *ReportService) GroupFinancialReport(groupID int64, categoryID int64) ([]dto.FinancialReportItem, error) {
	if groupID <= 0 {
		return nil, ErrInvalidInput
	}

	return s.reportRepository.FindGroupFinancialReport(groupID, categoryID)
}

func (s *ReportService) IncomeExpense(from *time.Time, to *time.Time) ([]dto.IncomeExpenseReportItem, error) {
	if invalidPeriod(from, to) {
		return nil, ErrInvalidInput
	}

	return s.reportRepository.FindIncomeExpense(from, to)
}

func (s *ReportService) CargoTypeShare(from *time.Time, to *time.Time) ([]dto.CargoTypeShareReportItem, error) {
	if invalidPeriod(from, to) {
		return nil, ErrInvalidInput
	}

	return s.reportRepository.FindCargoTypeShare(from, to)
}

func (s *ReportService) Profitability(from *time.Time, to *time.Time) (*dto.ProfitabilityReportResponse, error) {
	if invalidPeriod(from, to) {
		return nil, ErrInvalidInput
	}

	return s.reportRepository.FindProfitability(from, to)
}

func (s *ReportService) TouristCategoryRatio(
	from *time.Time,
	to *time.Time,
	restCategoryID int64,
	shopCategoryID int64,
) (*dto.TouristCategoryRatioReportResponse, error) {
	if invalidPeriod(from, to) {
		return nil, ErrInvalidInput
	}

	if restCategoryID <= 0 || shopCategoryID <= 0 || restCategoryID == shopCategoryID {
		return nil, ErrInvalidInput
	}

	return s.reportRepository.FindTouristCategoryRatio(from, to, restCategoryID, shopCategoryID)
}

func (s *ReportService) FlightTourists(flightID int64) ([]dto.FlightTouristReportItem, error) {
	if flightID <= 0 {
		return nil, ErrInvalidInput
	}

	return s.reportRepository.FindFlightTourists(flightID)
}

func invalidPeriod(from *time.Time, to *time.Time) bool {
	if from == nil || to == nil {
		return false
	}

	return from.After(*to)
}
