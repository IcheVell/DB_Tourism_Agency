package service

import (
	"errors"
	"strings"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"
)

var (
	ErrCargoTypeNotFound    = errors.New("cargo type not found")
	ErrCargoTypeInvalidName = errors.New("cargo type name is required")
)

type CargoTypeService struct {
	repo *repository.CargoTypeRepository
}

func NewCargoTypeService(repo *repository.CargoTypeRepository) *CargoTypeService {
	return &CargoTypeService{
		repo: repo,
	}
}

func (s *CargoTypeService) Create(req dto.CreateCargoTypeRequest) (*dto.CargoTypeResponse, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, ErrCargoTypeInvalidName
	}

	cargoType := &models.CargoType{
		Name: name,
	}

	if err := s.repo.Create(cargoType); err != nil {
		return nil, err
	}

	return toCargoTypeResponse(cargoType), nil
}

func (s *CargoTypeService) GetByID(id int64) (*dto.CargoTypeResponse, error) {
	cargoType, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if cargoType == nil {
		return nil, ErrCargoTypeNotFound
	}

	return toCargoTypeResponse(cargoType), nil
}

func (s *CargoTypeService) List(page int, pageSize int) (*dto.CargoTypeListResponse, error) {
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

	cargoTypes, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.CargoTypeResponse, 0, len(cargoTypes))
	for _, cargoType := range cargoTypes {
		items = append(items, *toCargoTypeResponse(&cargoType))
	}

	return &dto.CargoTypeListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *CargoTypeService) Update(id int64, req dto.UpdateCargoTypeRequest) (*dto.CargoTypeResponse, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, ErrCargoTypeInvalidName
	}

	cargoType, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if cargoType == nil {
		return nil, ErrCargoTypeNotFound
	}

	cargoType.Name = name

	if err := s.repo.Update(cargoType); err != nil {
		return nil, err
	}

	return toCargoTypeResponse(cargoType), nil
}

func (s *CargoTypeService) Delete(id int64) error {
	cargoType, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if cargoType == nil {
		return ErrCargoTypeNotFound
	}

	return s.repo.Delete(cargoType)
}

func toCargoTypeResponse(cargoType *models.CargoType) *dto.CargoTypeResponse {
	return &dto.CargoTypeResponse{
		ID:   cargoType.ID,
		Name: cargoType.Name,
	}
}
