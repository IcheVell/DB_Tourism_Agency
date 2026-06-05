package service

import (
	"errors"
	"strings"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"
)

var (
	ErrExcursionAgencyNotFound    = errors.New("excursion agency not found")
	ErrExcursionAgencyInvalidName = errors.New("excursion agency name is required")
)

type ExcursionAgencyService struct {
	repo *repository.ExcursionAgencyRepository
}

func NewExcursionAgencyService(repo *repository.ExcursionAgencyRepository) *ExcursionAgencyService {
	return &ExcursionAgencyService{
		repo: repo,
	}
}

func (s *ExcursionAgencyService) Create(req dto.CreateExcursionAgencyRequest) (*dto.ExcursionAgencyResponse, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, ErrExcursionAgencyInvalidName
	}

	agency := &models.ExcursionAgency{
		Name: name,
	}

	if err := s.repo.Create(agency); err != nil {
		return nil, err
	}

	return toExcursionAgencyResponse(agency), nil
}

func (s *ExcursionAgencyService) GetByID(id int64) (*dto.ExcursionAgencyResponse, error) {
	agency, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if agency == nil {
		return nil, ErrExcursionAgencyNotFound
	}

	return toExcursionAgencyResponse(agency), nil
}

func (s *ExcursionAgencyService) List(page int, pageSize int) (*dto.ExcursionAgencyListResponse, error) {
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

	agencies, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.ExcursionAgencyResponse, 0, len(agencies))
	for _, agency := range agencies {
		items = append(items, *toExcursionAgencyResponse(&agency))
	}

	return &dto.ExcursionAgencyListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *ExcursionAgencyService) Update(id int64, req dto.UpdateExcursionAgencyRequest) (*dto.ExcursionAgencyResponse, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, ErrExcursionAgencyInvalidName
	}

	agency, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if agency == nil {
		return nil, ErrExcursionAgencyNotFound
	}

	agency.Name = name

	if err := s.repo.Update(agency); err != nil {
		return nil, err
	}

	return toExcursionAgencyResponse(agency), nil
}

func (s *ExcursionAgencyService) Delete(id int64) error {
	agency, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if agency == nil {
		return ErrExcursionAgencyNotFound
	}

	return s.repo.Delete(agency)
}

func toExcursionAgencyResponse(agency *models.ExcursionAgency) *dto.ExcursionAgencyResponse {
	return &dto.ExcursionAgencyResponse{
		ID:   agency.ID,
		Name: agency.Name,
	}
}
