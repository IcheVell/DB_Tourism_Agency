package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type ExcursionScheduleRepository struct {
	db *gorm.DB
}

func NewExcursionScheduleRepository(db *gorm.DB) *ExcursionScheduleRepository {
	return &ExcursionScheduleRepository{
		db: db,
	}
}

func (r *ExcursionScheduleRepository) Create(schedule *models.ExcursionSchedule) error {
	return r.db.Create(schedule).Error
}

func (r *ExcursionScheduleRepository) FindByID(id int64) (*models.ExcursionSchedule, error) {
	var schedule models.ExcursionSchedule

	err := r.db.
		Where("id = ?", id).
		First(&schedule).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

func (r *ExcursionScheduleRepository) FindAll(limit int, offset int) ([]models.ExcursionSchedule, int64, error) {
	var schedules []models.ExcursionSchedule
	var total int64

	if err := r.db.Model(&models.ExcursionSchedule{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&schedules).
		Error

	if err != nil {
		return nil, 0, err
	}

	return schedules, total, nil
}

func (r *ExcursionScheduleRepository) Update(schedule *models.ExcursionSchedule) error {
	return r.db.Save(schedule).Error
}

func (r *ExcursionScheduleRepository) Delete(schedule *models.ExcursionSchedule) error {
	return r.db.Delete(schedule).Error
}
