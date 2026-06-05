package service

import (
	"errors"
	"strings"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"
)

var (
	ErrFlightTypeNotFound    = errors.New("flight type not found")
	ErrFlightTypeInvalidName = errors.New("flight type name is required")
)

type FlightTypeService struct {
	repo *repository.FlightTypeRepository
}

func NewFlightTypeService(repo *repository.FlightTypeRepository) *FlightTypeService {
	return &FlightTypeService{
		repo: repo,
	}
}

func (s *FlightTypeService) Create(req dto.CreateFlightTypeRequest) (*dto.FlightTypeResponse, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, ErrFlightTypeInvalidName
	}

	flightType := &models.FlightType{
		Name: name,
	}

	if err := s.repo.Create(flightType); err != nil {
		return nil, err
	}

	return toFlightTypeResponse(flightType), nil
}

func (s *FlightTypeService) GetByID(id int64) (*dto.FlightTypeResponse, error) {
	flightType, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if flightType == nil {
		return nil, ErrFlightTypeNotFound
	}

	return toFlightTypeResponse(flightType), nil
}

func (s *FlightTypeService) List(page int, pageSize int) (*dto.FlightTypeListResponse, error) {
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

	flightTypes, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.FlightTypeResponse, 0, len(flightTypes))
	for _, flightType := range flightTypes {
		items = append(items, *toFlightTypeResponse(&flightType))
	}

	return &dto.FlightTypeListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *FlightTypeService) Update(id int64, req dto.UpdateFlightTypeRequest) (*dto.FlightTypeResponse, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, ErrFlightTypeInvalidName
	}

	flightType, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if flightType == nil {
		return nil, ErrFlightTypeNotFound
	}

	flightType.Name = name

	if err := s.repo.Update(flightType); err != nil {
		return nil, err
	}

	return toFlightTypeResponse(flightType), nil
}

func (s *FlightTypeService) Delete(id int64) error {
	flightType, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if flightType == nil {
		return ErrFlightTypeNotFound
	}

	return s.repo.Delete(flightType)
}

func toFlightTypeResponse(flightType *models.FlightType) *dto.FlightTypeResponse {
	return &dto.FlightTypeResponse{
		ID:   flightType.ID,
		Name: flightType.Name,
	}
}
