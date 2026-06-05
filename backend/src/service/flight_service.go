package service

import (
	"errors"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"
)

var (
	ErrFlightNotFound            = errors.New("flight not found")
	ErrFlightInvalidCapacity     = errors.New("capacity must be positive")
	ErrFlightInvalidFlightDate   = errors.New("flight_date must be YYYY-MM-DD date")
	ErrFlightInvalidFlightTypeID = errors.New("flight_type_id must be positive")
)

type FlightService struct {
	repo *repository.FlightRepository
}

func NewFlightService(repo *repository.FlightRepository) *FlightService {
	return &FlightService{
		repo: repo,
	}
}

func (s *FlightService) Create(req dto.CreateFlightRequest) (*dto.FlightResponse, error) {
	flight, err := buildFlight(req.Capacity, req.FlightDate, req.FlightTypeID)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(flight); err != nil {
		return nil, err
	}

	return toFlightResponse(flight), nil
}

func (s *FlightService) GetByID(id int64) (*dto.FlightResponse, error) {
	flight, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if flight == nil {
		return nil, ErrFlightNotFound
	}

	return toFlightResponse(flight), nil
}

func (s *FlightService) List(page int, pageSize int) (*dto.FlightListResponse, error) {
	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 20
	}

	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	flights, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.FlightResponse, 0, len(flights))
	for _, flight := range flights {
		items = append(items, *toFlightResponse(&flight))
	}

	return &dto.FlightListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *FlightService) Update(id int64, req dto.UpdateFlightRequest) (*dto.FlightResponse, error) {
	flight, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if flight == nil {
		return nil, ErrFlightNotFound
	}

	updatedFlight, err := buildFlight(req.Capacity, req.FlightDate, req.FlightTypeID)
	if err != nil {
		return nil, err
	}

	flight.Capacity = updatedFlight.Capacity
	flight.FlightDate = updatedFlight.FlightDate
	flight.FlightTypeID = updatedFlight.FlightTypeID

	if err := s.repo.Update(flight); err != nil {
		return nil, err
	}

	return toFlightResponse(flight), nil
}

func (s *FlightService) Delete(id int64) error {
	flight, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if flight == nil {
		return ErrFlightNotFound
	}

	return s.repo.Delete(flight)
}

func buildFlight(capacity int, flightDateRaw string, flightTypeID int64) (*models.Flight, error) {
	if capacity <= 0 {
		return nil, ErrFlightInvalidCapacity
	}

	flightDate, err := parseDateYYYYMMDD(flightDateRaw)
	if err != nil {
		return nil, ErrFlightInvalidFlightDate
	}

	if flightTypeID <= 0 {
		return nil, ErrFlightInvalidFlightTypeID
	}

	return &models.Flight{
		Capacity:     capacity,
		FlightDate:   flightDate,
		FlightTypeID: flightTypeID,
	}, nil
}

func toFlightResponse(flight *models.Flight) *dto.FlightResponse {
	return &dto.FlightResponse{
		ID:           flight.ID,
		Capacity:     flight.Capacity,
		FlightDate:   flight.FlightDate.Format("2006-01-02"),
		FlightTypeID: flight.FlightTypeID,
	}
}
