package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type ExcursionBookingRepository struct {
	db *gorm.DB
}

func NewExcursionBookingRepository(db *gorm.DB) *ExcursionBookingRepository {
	return &ExcursionBookingRepository{
		db: db,
	}
}

func (r *ExcursionBookingRepository) Create(booking *models.ExcursionBooking) error {
	return r.db.Create(booking).Error
}

func (r *ExcursionBookingRepository) FindByID(id int64) (*models.ExcursionBooking, error) {
	var booking models.ExcursionBooking

	err := r.db.
		Where("id = ?", id).
		First(&booking).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &booking, nil
}

func (r *ExcursionBookingRepository) FindAll(limit int, offset int) ([]models.ExcursionBooking, int64, error) {
	var bookings []models.ExcursionBooking
	var total int64

	if err := r.db.Model(&models.ExcursionBooking{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&bookings).
		Error

	if err != nil {
		return nil, 0, err
	}

	return bookings, total, nil
}

func (r *ExcursionBookingRepository) Update(booking *models.ExcursionBooking) error {
	return r.db.Save(booking).Error
}

func (r *ExcursionBookingRepository) Delete(booking *models.ExcursionBooking) error {
	return r.db.Delete(booking).Error
}
