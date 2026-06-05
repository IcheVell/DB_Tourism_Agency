package service

import (
	"errors"
	"strings"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"
)

var (
	ErrFinancialCategoryNotFound             = errors.New("financial category not found")
	ErrFinancialCategoryInvalidName          = errors.New("financial category name is required")
	ErrFinancialCategoryInvalidOperationType = errors.New("financial category operation type must be income or expense")
)

type FinancialCategoryService struct {
	repo *repository.FinancialCategoryRepository
}

func NewFinancialCategoryService(repo *repository.FinancialCategoryRepository) *FinancialCategoryService {
	return &FinancialCategoryService{
		repo: repo,
	}
}

func (s *FinancialCategoryService) Create(req dto.CreateFinancialCategoryRequest) (*dto.FinancialCategoryResponse, error) {
	name := strings.TrimSpace(req.Name)
	operationType := strings.TrimSpace(req.OperationType)

	if name == "" {
		return nil, ErrFinancialCategoryInvalidName
	}

	if operationType != "income" && operationType != "expense" {
		return nil, ErrFinancialCategoryInvalidOperationType
	}

	category := &models.FinancialCategory{
		Name:          name,
		OperationType: operationType,
	}

	if err := s.repo.Create(category); err != nil {
		return nil, err
	}

	return toFinancialCategoryResponse(category), nil
}

func (s *FinancialCategoryService) GetByID(id int64) (*dto.FinancialCategoryResponse, error) {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, ErrFinancialCategoryNotFound
	}

	return toFinancialCategoryResponse(category), nil
}

func (s *FinancialCategoryService) List(page int, pageSize int) (*dto.FinancialCategoryListResponse, error) {
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

	items := make([]dto.FinancialCategoryResponse, 0, len(categories))
	for _, category := range categories {
		items = append(items, *toFinancialCategoryResponse(&category))
	}

	return &dto.FinancialCategoryListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *FinancialCategoryService) Update(id int64, req dto.UpdateFinancialCategoryRequest) (*dto.FinancialCategoryResponse, error) {
	name := strings.TrimSpace(req.Name)
	operationType := strings.TrimSpace(req.OperationType)

	if name == "" {
		return nil, ErrFinancialCategoryInvalidName
	}

	if operationType != "income" && operationType != "expense" {
		return nil, ErrFinancialCategoryInvalidOperationType
	}

	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, ErrFinancialCategoryNotFound
	}

	category.Name = name
	category.OperationType = operationType

	if err := s.repo.Update(category); err != nil {
		return nil, err
	}

	return toFinancialCategoryResponse(category), nil
}

func (s *FinancialCategoryService) Delete(id int64) error {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if category == nil {
		return ErrFinancialCategoryNotFound
	}

	return s.repo.Delete(category)
}

func toFinancialCategoryResponse(category *models.FinancialCategory) *dto.FinancialCategoryResponse {
	return &dto.FinancialCategoryResponse{
		ID:            category.ID,
		Name:          category.Name,
		OperationType: category.OperationType,
	}
}
