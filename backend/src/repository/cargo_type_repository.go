package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type CargoTypeRepository struct {
	db *gorm.DB
}

func NewCargoTypeRepository(db *gorm.DB) *CargoTypeRepository {
	return &CargoTypeRepository{
		db: db,
	}
}

func (r *CargoTypeRepository) Create(cargoType *models.CargoType) error {
	return r.db.Create(cargoType).Error
}

func (r *CargoTypeRepository) FindByID(id int64) (*models.CargoType, error) {
	var cargoType models.CargoType

	err := r.db.
		Where("id = ?", id).
		First(&cargoType).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &cargoType, nil
}

func (r *CargoTypeRepository) FindAll(limit int, offset int) ([]models.CargoType, int64, error) {
	var cargoTypes []models.CargoType
	var total int64

	if err := r.db.Model(&models.CargoType{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&cargoTypes).
		Error

	if err != nil {
		return nil, 0, err
	}

	return cargoTypes, total, nil
}

func (r *CargoTypeRepository) Update(cargoType *models.CargoType) error {
	return r.db.Save(cargoType).Error
}

func (r *CargoTypeRepository) Delete(cargoType *models.CargoType) error {
	return r.db.Delete(cargoType).Error
}
