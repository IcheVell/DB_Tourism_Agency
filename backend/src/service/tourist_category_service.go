package service

import (
	"errors"
	"strings"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"
)

var (
	ErrTouristCategoryNotFound    = errors.New("tourist category not found")
	ErrTouristCategoryInvalidName = errors.New("tourist category name is required")
)

type TouristCategoryService struct {
	repo *repository.TouristCategoryRepository
}

func NewTouristCategoryService(repo *repository.TouristCategoryRepository) *TouristCategoryService {
	return &TouristCategoryService{
		repo: repo,
	}
}

func (s *TouristCategoryService) Create(req dto.CreateTouristCategoryRequest) (*dto.TouristCategoryResponse, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, ErrTouristCategoryInvalidName
	}

	category := &models.TouristCategory{
		Name: name,
	}

	if err := s.repo.Create(category); err != nil {
		return nil, err
	}

	return toTouristCategoryResponse(category), nil
}

func (s *TouristCategoryService) GetByID(id int64) (*dto.TouristCategoryResponse, error) {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, ErrTouristCategoryNotFound
	}

	return toTouristCategoryResponse(category), nil
}

func (s *TouristCategoryService) List(page int, pageSize int) (*dto.TouristCategoryListResponse, error) {
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

	categories, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.TouristCategoryResponse, 0, len(categories))
	for _, category := range categories {
		items = append(items, *toTouristCategoryResponse(&category))
	}

	return &dto.TouristCategoryListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *TouristCategoryService) Update(id int64, req dto.UpdateTouristCategoryRequest) (*dto.TouristCategoryResponse, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, ErrTouristCategoryInvalidName
	}

	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, ErrTouristCategoryNotFound
	}

	category.Name = name

	if err := s.repo.Update(category); err != nil {
		return nil, err
	}

	return toTouristCategoryResponse(category), nil
}

func (s *TouristCategoryService) Delete(id int64) error {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if category == nil {
		return ErrTouristCategoryNotFound
	}

	return s.repo.Delete(category)
}

func toTouristCategoryResponse(category *models.TouristCategory) *dto.TouristCategoryResponse {
	return &dto.TouristCategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}
