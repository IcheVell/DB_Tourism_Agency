package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type TouristRepository struct {
	db *gorm.DB
}

func NewTouristRepository(db *gorm.DB) *TouristRepository {
	return &TouristRepository{
		db: db,
	}
}

func (r *TouristRepository) Create(tourist *models.Tourist) error {
	return r.db.Create(tourist).Error
}

func (r *TouristRepository) FindByID(id int64) (*models.Tourist, error) {
	var tourist models.Tourist

	err := r.db.
		Where("id = ?", id).
		First(&tourist).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &tourist, nil
}

func (r *TouristRepository) FindAll(limit int, offset int) ([]models.Tourist, int64, error) {
	var tourists []models.Tourist
	var total int64

	if err := r.db.Model(&models.Tourist{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&tourists).
		Error

	if err != nil {
		return nil, 0, err
	}

	return tourists, total, nil
}

func (r *TouristRepository) Update(tourist *models.Tourist) error {
	return r.db.Save(tourist).Error
}

func (r *TouristRepository) Delete(tourist *models.Tourist) error {
	return r.db.Delete(tourist).Error
}
