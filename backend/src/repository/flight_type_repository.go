package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type FlightTypeRepository struct {
	db *gorm.DB
}

func NewFlightTypeRepository(db *gorm.DB) *FlightTypeRepository {
	return &FlightTypeRepository{
		db: db,
	}
}

func (r *FlightTypeRepository) Create(flightType *models.FlightType) error {
	return r.db.Create(flightType).Error
}

func (r *FlightTypeRepository) FindByID(id int64) (*models.FlightType, error) {
	var flightType models.FlightType

	err := r.db.
		Where("id = ?", id).
		First(&flightType).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &flightType, nil
}

func (r *FlightTypeRepository) FindAll(limit int, offset int) ([]models.FlightType, int64, error) {
	var flightTypes []models.FlightType
	var total int64

	if err := r.db.Model(&models.FlightType{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&flightTypes).
		Error

	if err != nil {
		return nil, 0, err
	}

	return flightTypes, total, nil
}

func (r *FlightTypeRepository) Update(flightType *models.FlightType) error {
	return r.db.Save(flightType).Error
}

func (r *FlightTypeRepository) Delete(flightType *models.FlightType) error {
	return r.db.Delete(flightType).Error
}
