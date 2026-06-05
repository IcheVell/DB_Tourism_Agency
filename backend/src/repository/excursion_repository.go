package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type ExcursionRepository struct {
	db *gorm.DB
}

func NewExcursionRepository(db *gorm.DB) *ExcursionRepository {
	return &ExcursionRepository{
		db: db,
	}
}

func (r *ExcursionRepository) Create(excursion *models.Excursion) error {
	return r.db.Create(excursion).Error
}

func (r *ExcursionRepository) FindByID(id int64) (*models.Excursion, error) {
	var excursion models.Excursion

	err := r.db.
		Where("id = ?", id).
		First(&excursion).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &excursion, nil
}

func (r *ExcursionRepository) FindAll(limit int, offset int) ([]models.Excursion, int64, error) {
	var excursions []models.Excursion
	var total int64

	if err := r.db.Model(&models.Excursion{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&excursions).
		Error

	if err != nil {
		return nil, 0, err
	}

	return excursions, total, nil
}

func (r *ExcursionRepository) Update(excursion *models.Excursion) error {
	return r.db.Save(excursion).Error
}

func (r *ExcursionRepository) Delete(excursion *models.Excursion) error {
	return r.db.Delete(excursion).Error
}
