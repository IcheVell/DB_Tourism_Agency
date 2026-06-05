package service

import (
	"errors"
	"strings"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"
)

var (
	ErrExcursionNotFound    = errors.New("excursion not found")
	ErrExcursionInvalidName = errors.New("excursion name is required")
)

type ExcursionService struct {
	repo *repository.ExcursionRepository
}

func NewExcursionService(repo *repository.ExcursionRepository) *ExcursionService {
	return &ExcursionService{
		repo: repo,
	}
}

func (s *ExcursionService) Create(req dto.CreateExcursionRequest) (*dto.ExcursionResponse, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, ErrExcursionInvalidName
	}

	description := normalizeOptionalString(req.Description)

	excursion := &models.Excursion{
		Name:        name,
		Description: description,
	}

	if err := s.repo.Create(excursion); err != nil {
		return nil, err
	}

	return toExcursionResponse(excursion), nil
}

func (s *ExcursionService) GetByID(id int64) (*dto.ExcursionResponse, error) {
	excursion, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if excursion == nil {
		return nil, ErrExcursionNotFound
	}

	return toExcursionResponse(excursion), nil
}

func (s *ExcursionService) List(page int, pageSize int) (*dto.ExcursionListResponse, error) {
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

	excursions, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.ExcursionResponse, 0, len(excursions))
	for _, excursion := range excursions {
		items = append(items, *toExcursionResponse(&excursion))
	}

	return &dto.ExcursionListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *ExcursionService) Update(id int64, req dto.UpdateExcursionRequest) (*dto.ExcursionResponse, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, ErrExcursionInvalidName
	}

	excursion, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if excursion == nil {
		return nil, ErrExcursionNotFound
	}

	excursion.Name = name
	excursion.Description = normalizeOptionalString(req.Description)

	if err := s.repo.Update(excursion); err != nil {
		return nil, err
	}

	return toExcursionResponse(excursion), nil
}

func (s *ExcursionService) Delete(id int64) error {
	excursion, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if excursion == nil {
		return ErrExcursionNotFound
	}

	return s.repo.Delete(excursion)
}

func toExcursionResponse(excursion *models.Excursion) *dto.ExcursionResponse {
	return &dto.ExcursionResponse{
		ID:          excursion.ID,
		Name:        excursion.Name,
		Description: excursion.Description,
	}
}
