package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type HotelRoomRepository struct {
	db *gorm.DB
}

func NewHotelRoomRepository(db *gorm.DB) *HotelRoomRepository {
	return &HotelRoomRepository{
		db: db,
	}
}

func (r *HotelRoomRepository) Create(room *models.HotelRoom) error {
	return r.db.Create(room).Error
}

func (r *HotelRoomRepository) FindByID(id int64) (*models.HotelRoom, error) {
	var room models.HotelRoom

	err := r.db.
		Where("id = ?", id).
		First(&room).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (r *HotelRoomRepository) FindAll(limit int, offset int) ([]models.HotelRoom, int64, error) {
	var rooms []models.HotelRoom
	var total int64

	if err := r.db.Model(&models.HotelRoom{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&rooms).
		Error

	if err != nil {
		return nil, 0, err
	}

	return rooms, total, nil
}

func (r *HotelRoomRepository) Update(room *models.HotelRoom) error {
	return r.db.Save(room).Error
}

func (r *HotelRoomRepository) Delete(room *models.HotelRoom) error {
	return r.db.Delete(room).Error
}
