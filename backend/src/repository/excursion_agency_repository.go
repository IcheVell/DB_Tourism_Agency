package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type ExcursionAgencyRepository struct {
	db *gorm.DB
}

func NewExcursionAgencyRepository(db *gorm.DB) *ExcursionAgencyRepository {
	return &ExcursionAgencyRepository{
		db: db,
	}
}

func (r *ExcursionAgencyRepository) Create(agency *models.ExcursionAgency) error {
	return r.db.Create(agency).Error
}

func (r *ExcursionAgencyRepository) FindByID(id int64) (*models.ExcursionAgency, error) {
	var agency models.ExcursionAgency

	err := r.db.
		Where("id = ?", id).
		First(&agency).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &agency, nil
}

func (r *ExcursionAgencyRepository) FindAll(limit int, offset int) ([]models.ExcursionAgency, int64, error) {
	var agencies []models.ExcursionAgency
	var total int64

	if err := r.db.Model(&models.ExcursionAgency{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&agencies).
		Error

	if err != nil {
		return nil, 0, err
	}

	return agencies, total, nil
}

func (r *ExcursionAgencyRepository) Update(agency *models.ExcursionAgency) error {
	return r.db.Save(agency).Error
}

func (r *ExcursionAgencyRepository) Delete(agency *models.ExcursionAgency) error {
	return r.db.Delete(agency).Error
}
