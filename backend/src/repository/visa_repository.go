package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type VisaRepository struct {
	db *gorm.DB
}

func NewVisaRepository(db *gorm.DB) *VisaRepository {
	return &VisaRepository{
		db: db,
	}
}

func (r *VisaRepository) Create(visa *models.Visa) error {
	return r.db.Create(visa).Error
}

func (r *VisaRepository) FindByID(id int64) (*models.Visa, error) {
	var visa models.Visa

	err := r.db.
		Where("id = ?", id).
		First(&visa).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &visa, nil
}

func (r *VisaRepository) FindAll(limit int, offset int) ([]models.Visa, int64, error) {
	var visas []models.Visa
	var total int64

	if err := r.db.Model(&models.Visa{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&visas).
		Error

	if err != nil {
		return nil, 0, err
	}

	return visas, total, nil
}

func (r *VisaRepository) Update(visa *models.Visa) error {
	return r.db.Save(visa).Error
}

func (r *VisaRepository) Delete(visa *models.Visa) error {
	return r.db.Delete(visa).Error
}
