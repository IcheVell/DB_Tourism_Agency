package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type UserRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) *UserRoleRepository {
	return &UserRoleRepository{
		db: db,
	}
}

func (r *UserRoleRepository) Create(userRole *models.UserRole) error {
	return r.db.Create(userRole).Error
}

func (r *UserRoleRepository) FindByIDs(userID int64, roleID int64) (*models.UserRole, error) {
	var userRole models.UserRole

	err := r.db.
		Preload("User").
		Preload("Role").
		Where("user_id = ? AND role_id = ?", userID, roleID).
		First(&userRole).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &userRole, nil
}

func (r *UserRoleRepository) FindByUserID(userID int64) (*models.UserRole, error) {
	var userRole models.UserRole

	err := r.db.
		Preload("User").
		Preload("Role").
		Where("user_id = ?", userID).
		First(&userRole).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &userRole, nil
}

func (r *UserRoleRepository) FindAll(limit int, offset int) ([]models.UserRole, int64, error) {
	var userRoles []models.UserRole
	var total int64

	if err := r.db.Model(&models.UserRole{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Preload("User").
		Preload("Role").
		Order("user_id ASC, role_id ASC").
		Limit(limit).
		Offset(offset).
		Find(&userRoles).
		Error

	if err != nil {
		return nil, 0, err
	}

	return userRoles, total, nil
}

func (r *UserRoleRepository) FindAllByUserID(userID int64, limit int, offset int) ([]models.UserRole, int64, error) {
	var userRoles []models.UserRole
	var total int64

	query := r.db.
		Model(&models.UserRole{}).
		Where("user_id = ?", userID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Preload("User").
		Preload("Role").
		Where("user_id = ?", userID).
		Order("role_id ASC").
		Limit(limit).
		Offset(offset).
		Find(&userRoles).
		Error

	if err != nil {
		return nil, 0, err
	}

	return userRoles, total, nil
}

func (r *UserRoleRepository) Update(
	oldUserID int64,
	oldRoleID int64,
	userRole *models.UserRole,
) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var existing models.UserRole

		err := tx.
			Where("user_id = ? AND role_id = ?", oldUserID, oldRoleID).
			First(&existing).
			Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}

		if err != nil {
			return err
		}

		if oldUserID == userRole.UserID && oldRoleID == userRole.RoleID {
			return nil
		}

		if err := tx.
			Where("user_id = ? AND role_id = ?", oldUserID, oldRoleID).
			Delete(&models.UserRole{}).
			Error; err != nil {
			return err
		}

		return tx.Create(userRole).Error
	})
}

func (r *UserRoleRepository) Delete(userRole *models.UserRole) error {
	return r.db.
		Where("user_id = ? AND role_id = ?", userRole.UserID, userRole.RoleID).
		Delete(&models.UserRole{}).
		Error
}

func (r *UserRoleRepository) ExistsByIDs(userID int64, roleID int64) (bool, error) {
	var count int64

	err := r.db.
		Model(&models.UserRole{}).
		Where("user_id = ? AND role_id = ?", userID, roleID).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *UserRoleRepository) UserHasRole(userID int64) (bool, error) {
	var count int64

	err := r.db.
		Model(&models.UserRole{}).
		Where("user_id = ?", userID).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *UserRoleRepository) UserExists(userID int64) (bool, error) {
	var count int64

	err := r.db.
		Model(&models.User{}).
		Where("id = ?", userID).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *UserRoleRepository) RoleExists(roleID int64) (bool, error) {
	var count int64

	err := r.db.
		Model(&models.Role{}).
		Where("id = ?", roleID).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
