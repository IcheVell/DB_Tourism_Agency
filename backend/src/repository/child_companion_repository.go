package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type ChildCompanionRepository struct {
	db *gorm.DB
}

func NewChildCompanionRepository(db *gorm.DB) *ChildCompanionRepository {
	return &ChildCompanionRepository{
		db: db,
	}
}

func (r *ChildCompanionRepository) Create(companion *models.ChildCompanion) error {
	return r.db.Create(companion).Error
}

func (r *ChildCompanionRepository) FindByIDs(adultGroupMemberID int64, childGroupMemberID int64) (*models.ChildCompanion, error) {
	var companion models.ChildCompanion

	err := r.db.
		Where(
			"adult_group_member_id = ? AND child_group_member_id = ?",
			adultGroupMemberID,
			childGroupMemberID,
		).
		First(&companion).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &companion, nil
}

func (r *ChildCompanionRepository) FindAll(limit int, offset int) ([]models.ChildCompanion, int64, error) {
	var companions []models.ChildCompanion
	var total int64

	if err := r.db.Model(&models.ChildCompanion{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("adult_group_member_id ASC, child_group_member_id ASC").
		Limit(limit).
		Offset(offset).
		Find(&companions).
		Error

	if err != nil {
		return nil, 0, err
	}

	return companions, total, nil
}

func (r *ChildCompanionRepository) Delete(companion *models.ChildCompanion) error {
	return r.db.Delete(companion).Error
}

func (r *ChildCompanionRepository) Replace(
	oldCompanion *models.ChildCompanion,
	newCompanion *models.ChildCompanion,
) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(oldCompanion).Error; err != nil {
			return err
		}

		if err := tx.Create(newCompanion).Error; err != nil {
			return err
		}

		return nil
	})
}
