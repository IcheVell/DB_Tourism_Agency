package repository

import (
	"TouristAgencyApp/src/models"
	"errors"

	"gorm.io/gorm"
)

type FinancialCategoryRepository struct {
	db *gorm.DB
}

func NewFinancialCategoryRepository(db *gorm.DB) *FinancialCategoryRepository {
	return &FinancialCategoryRepository{
		db: db,
	}
}

func (r *FinancialCategoryRepository) Create(category *models.FinancialCategory) error {
	return r.db.Create(category).Error
}

func (r *FinancialCategoryRepository) FindByID(id int64) (*models.FinancialCategory, error) {
	var category models.FinancialCategory

	err := r.db.
		Where("id = ?", id).
		First(&category).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *FinancialCategoryRepository) FindAll(limit int, offset int) ([]models.FinancialCategory, int64, error) {
	var categories []models.FinancialCategory
	var total int64

	if err := r.db.Model(&models.FinancialCategory{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&categories).
		Error

	if err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

func (r *FinancialCategoryRepository) Update(category *models.FinancialCategory) error {
	return r.db.Save(category).Error
}

func (r *FinancialCategoryRepository) Delete(category *models.FinancialCategory) error {
	return r.db.Delete(category).Error
}
