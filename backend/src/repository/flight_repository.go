package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type FlightRepository struct {
	db *gorm.DB
}

func NewFlightRepository(db *gorm.DB) *FlightRepository {
	return &FlightRepository{
		db: db,
	}
}

func (r *FlightRepository) Create(flight *models.Flight) error {
	return r.db.Create(flight).Error
}

func (r *FlightRepository) FindByID(id int64) (*models.Flight, error) {
	var flight models.Flight

	err := r.db.
		Where("id = ?", id).
		First(&flight).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &flight, nil
}

func (r *FlightRepository) FindAll(limit int, offset int) ([]models.Flight, int64, error) {
	var flights []models.Flight
	var total int64

	if err := r.db.Model(&models.Flight{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&flights).
		Error

	if err != nil {
		return nil, 0, err
	}

	return flights, total, nil
}

func (r *FlightRepository) Update(flight *models.Flight) error {
	return r.db.Save(flight).Error
}

func (r *FlightRepository) Delete(flight *models.Flight) error {
	return r.db.Delete(flight).Error
}
