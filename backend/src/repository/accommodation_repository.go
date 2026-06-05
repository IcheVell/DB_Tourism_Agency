package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type AccommodationRepository struct {
	db *gorm.DB
}

func NewAccommodationRepository(db *gorm.DB) *AccommodationRepository {
	return &AccommodationRepository{
		db: db,
	}
}

func (r *AccommodationRepository) Create(accommodation *models.Accommodation) error {
	return r.db.Create(accommodation).Error
}

func (r *AccommodationRepository) FindByID(id int64) (*models.Accommodation, error) {
	var accommodation models.Accommodation

	err := r.db.
		Where("id = ?", id).
		First(&accommodation).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &accommodation, nil
}

func (r *AccommodationRepository) FindAll(limit int, offset int) ([]models.Accommodation, int64, error) {
	var accommodations []models.Accommodation
	var total int64

	if err := r.db.Model(&models.Accommodation{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&accommodations).
		Error

	if err != nil {
		return nil, 0, err
	}

	return accommodations, total, nil
}

func (r *AccommodationRepository) Update(accommodation *models.Accommodation) error {
	return r.db.Save(accommodation).Error
}

func (r *AccommodationRepository) Delete(accommodation *models.Accommodation) error {
	return r.db.Delete(accommodation).Error
}
