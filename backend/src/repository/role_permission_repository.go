package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type RolePermissionRepository struct {
	db *gorm.DB
}

func NewRolePermissionRepository(db *gorm.DB) *RolePermissionRepository {
	return &RolePermissionRepository{
		db: db,
	}
}

func (r *RolePermissionRepository) Create(rolePermission *models.RolePermission) error {
	return r.db.Create(rolePermission).Error
}

func (r *RolePermissionRepository) FindByIDs(roleID int64, permissionID int64) (*models.RolePermission, error) {
	var rolePermission models.RolePermission

	err := r.db.
		Preload("Role").
		Preload("Permission").
		Where("role_id = ? AND permission_id = ?", roleID, permissionID).
		First(&rolePermission).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &rolePermission, nil
}

func (r *RolePermissionRepository) FindAll(limit int, offset int) ([]models.RolePermission, int64, error) {
	var rolePermissions []models.RolePermission
	var total int64

	if err := r.db.Model(&models.RolePermission{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Preload("Role").
		Preload("Permission").
		Order("role_id ASC, permission_id ASC").
		Limit(limit).
		Offset(offset).
		Find(&rolePermissions).
		Error

	if err != nil {
		return nil, 0, err
	}

	return rolePermissions, total, nil
}

func (r *RolePermissionRepository) FindByRoleID(roleID int64, limit int, offset int) ([]models.RolePermission, int64, error) {
	var rolePermissions []models.RolePermission
	var total int64

	query := r.db.
		Model(&models.RolePermission{}).
		Where("role_id = ?", roleID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Preload("Role").
		Preload("Permission").
		Where("role_id = ?", roleID).
		Order("permission_id ASC").
		Limit(limit).
		Offset(offset).
		Find(&rolePermissions).
		Error

	if err != nil {
		return nil, 0, err
	}

	return rolePermissions, total, nil
}

func (r *RolePermissionRepository) Update(
	oldRoleID int64,
	oldPermissionID int64,
	rolePermission *models.RolePermission,
) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var existing models.RolePermission

		err := tx.
			Where("role_id = ? AND permission_id = ?", oldRoleID, oldPermissionID).
			First(&existing).
			Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}

		if err != nil {
			return err
		}

		if oldRoleID == rolePermission.RoleID && oldPermissionID == rolePermission.PermissionID {
			return nil
		}

		if err := tx.
			Where("role_id = ? AND permission_id = ?", oldRoleID, oldPermissionID).
			Delete(&models.RolePermission{}).
			Error; err != nil {
			return err
		}

		return tx.Create(rolePermission).Error
	})
}

func (r *RolePermissionRepository) Delete(rolePermission *models.RolePermission) error {
	return r.db.
		Where("role_id = ? AND permission_id = ?", rolePermission.RoleID, rolePermission.PermissionID).
		Delete(&models.RolePermission{}).
		Error
}

func (r *RolePermissionRepository) ExistsByIDs(roleID int64, permissionID int64) (bool, error) {
	var count int64

	err := r.db.
		Model(&models.RolePermission{}).
		Where("role_id = ? AND permission_id = ?", roleID, permissionID).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *RolePermissionRepository) RoleExists(roleID int64) (bool, error) {
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

func (r *RolePermissionRepository) PermissionExists(permissionID int64) (bool, error) {
	var count int64

	err := r.db.
		Model(&models.Permission{}).
		Where("id = ?", permissionID).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
