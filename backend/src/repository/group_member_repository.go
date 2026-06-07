package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type GroupMemberRepository struct {
	db *gorm.DB
}

func NewGroupMemberRepository(db *gorm.DB) *GroupMemberRepository {
	return &GroupMemberRepository{
		db: db,
	}
}

func (r *GroupMemberRepository) FindTouristDesiredHotelID(touristID int64) (*int64, error) {
	var desiredHotelID int64

	err := r.db.Raw(`
		SELECT COALESCE(desired_hotel_id, 0)
		FROM tourists
		WHERE id = ?
	`, touristID).Scan(&desiredHotelID).Error

	if err != nil {
		return nil, err
	}

	if desiredHotelID == 0 {
		return nil, nil
	}

	return &desiredHotelID, nil
}

func (r *GroupMemberRepository) Create(groupMember *models.GroupMember) error {
	return r.db.Create(groupMember).Error
}

func (r *GroupMemberRepository) FindByID(id int64) (*models.GroupMember, error) {
	var groupMember models.GroupMember

	err := r.db.
		Where("id = ?", id).
		First(&groupMember).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &groupMember, nil
}

func (r *GroupMemberRepository) FindAll(limit int, offset int) ([]models.GroupMember, int64, error) {
	var groupMembers []models.GroupMember
	var total int64

	if err := r.db.Model(&models.GroupMember{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&groupMembers).
		Error

	if err != nil {
		return nil, 0, err
	}

	return groupMembers, total, nil
}

func (r *GroupMemberRepository) Update(groupMember *models.GroupMember) error {
	return r.db.Save(groupMember).Error
}

func (r *GroupMemberRepository) Delete(groupMember *models.GroupMember) error {
	return r.db.Delete(groupMember).Error
}
