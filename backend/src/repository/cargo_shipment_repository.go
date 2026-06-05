package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type CargoShipmentRepository struct {
	db *gorm.DB
}

func NewCargoShipmentRepository(db *gorm.DB) *CargoShipmentRepository {
	return &CargoShipmentRepository{
		db: db,
	}
}

func (r *CargoShipmentRepository) Create(shipment *models.CargoShipment) error {
	return r.db.Create(shipment).Error
}

func (r *CargoShipmentRepository) FindByID(id int64) (*models.CargoShipment, error) {
	var shipment models.CargoShipment

	err := r.db.
		Where("id = ?", id).
		First(&shipment).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &shipment, nil
}

func (r *CargoShipmentRepository) FindAll(limit int, offset int) ([]models.CargoShipment, int64, error) {
	var shipments []models.CargoShipment
	var total int64

	if err := r.db.Model(&models.CargoShipment{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&shipments).
		Error

	if err != nil {
		return nil, 0, err
	}

	return shipments, total, nil
}

func (r *CargoShipmentRepository) Update(shipment *models.CargoShipment) error {
	return r.db.Save(shipment).Error
}

func (r *CargoShipmentRepository) Delete(shipment *models.CargoShipment) error {
	return r.db.Delete(shipment).Error
}
