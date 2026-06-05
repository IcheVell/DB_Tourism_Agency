package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type CargoItemRepository struct {
	db *gorm.DB
}

func NewCargoItemRepository(db *gorm.DB) *CargoItemRepository {
	return &CargoItemRepository{
		db: db,
	}
}

func (r *CargoItemRepository) Create(item *models.CargoItem) error {
	return r.db.Create(item).Error
}

func (r *CargoItemRepository) FindByID(id int64) (*models.CargoItem, error) {
	var item models.CargoItem

	err := r.db.
		Where("id = ?", id).
		First(&item).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *CargoItemRepository) FindAll(limit int, offset int) ([]models.CargoItem, int64, error) {
	var items []models.CargoItem
	var total int64

	if err := r.db.Model(&models.CargoItem{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&items).
		Error

	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *CargoItemRepository) Update(item *models.CargoItem) error {
	return r.db.Save(item).Error
}

func (r *CargoItemRepository) Delete(item *models.CargoItem) error {
	return r.db.Delete(item).Error
}
