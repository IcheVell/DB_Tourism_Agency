package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type HotelRepository struct {
	db *gorm.DB
}

func NewHotelRepository(db *gorm.DB) *HotelRepository {
	return &HotelRepository{
		db: db,
	}
}

func (r *HotelRepository) Create(hotel *models.Hotel) error {
	return r.db.Create(hotel).Error
}

func (r *HotelRepository) FindByID(id int64) (*models.Hotel, error) {
	var hotel models.Hotel

	err := r.db.
		Where("id = ?", id).
		First(&hotel).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &hotel, nil
}

func (r *HotelRepository) FindAll(limit int, offset int) ([]models.Hotel, int64, error) {
	var hotels []models.Hotel
	var total int64

	if err := r.db.Model(&models.Hotel{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&hotels).
		Error

	if err != nil {
		return nil, 0, err
	}

	return hotels, total, nil
}

func (r *HotelRepository) Update(hotel *models.Hotel) error {
	return r.db.Save(hotel).Error
}

func (r *HotelRepository) Delete(hotel *models.Hotel) error {
	return r.db.Delete(hotel).Error
}
