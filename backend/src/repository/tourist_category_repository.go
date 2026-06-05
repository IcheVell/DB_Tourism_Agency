package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type TouristCategoryRepository struct {
	db *gorm.DB
}

func NewTouristCategoryRepository(db *gorm.DB) *TouristCategoryRepository {
	return &TouristCategoryRepository{
		db: db,
	}
}

func (r *TouristCategoryRepository) Create(category *models.TouristCategory) error {
	return r.db.Create(category).Error
}

func (r *TouristCategoryRepository) FindByID(id int64) (*models.TouristCategory, error) {
	var category models.TouristCategory

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

func (r *TouristCategoryRepository) FindAll(limit int, offset int) ([]models.TouristCategory, int64, error) {
	var categories []models.TouristCategory
	var total int64

	if err := r.db.Model(&models.TouristCategory{}).Count(&total).Error; err != nil {
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

func (r *TouristCategoryRepository) Update(category *models.TouristCategory) error {
	return r.db.Save(category).Error
}

func (r *TouristCategoryRepository) Delete(category *models.TouristCategory) error {
	return r.db.Delete(category).Error
}
