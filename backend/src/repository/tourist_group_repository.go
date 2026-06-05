package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type TouristGroupRepository struct {
	db *gorm.DB
}

func NewTouristGroupRepository(db *gorm.DB) *TouristGroupRepository {
	return &TouristGroupRepository{
		db: db,
	}
}

func (r *TouristGroupRepository) Create(group *models.TouristGroup) error {
	return r.db.Create(group).Error
}

func (r *TouristGroupRepository) FindByID(id int64) (*models.TouristGroup, error) {
	var group models.TouristGroup

	err := r.db.
		Where("id = ?", id).
		First(&group).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (r *TouristGroupRepository) FindAll(limit int, offset int) ([]models.TouristGroup, int64, error) {
	var groups []models.TouristGroup
	var total int64

	if err := r.db.Model(&models.TouristGroup{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&groups).
		Error

	if err != nil {
		return nil, 0, err
	}

	return groups, total, nil
}

func (r *TouristGroupRepository) Update(group *models.TouristGroup) error {
	return r.db.Save(group).Error
}

func (r *TouristGroupRepository) Delete(group *models.TouristGroup) error {
	return r.db.Delete(group).Error
}
