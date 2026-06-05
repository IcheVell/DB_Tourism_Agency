package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) Create(role *models.Role, permissionIDs []int64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(role).Error; err != nil {
			return err
		}

		if len(permissionIDs) == 0 {
			return nil
		}

		rolePermissions := make([]models.RolePermission, 0, len(permissionIDs))

		for _, permissionID := range permissionIDs {
			rolePermissions = append(rolePermissions, models.RolePermission{
				RoleID:       role.ID,
				PermissionID: permissionID,
			})
		}

		return tx.Create(&rolePermissions).Error
	})
}

func (r *RoleRepository) FindByID(id int64) (*models.Role, error) {
	var role models.Role

	err := r.db.
		Where("id = ?", id).
		First(&role).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *RoleRepository) FindAll(limit int, offset int) ([]models.Role, int64, error) {
	var roles []models.Role
	var total int64

	if err := r.db.Model(&models.Role{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&roles).
		Error

	if err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}

func (r *RoleRepository) Update(role *models.Role, permissionIDs *[]int64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(role).Error; err != nil {
			return err
		}

		if permissionIDs == nil {
			return nil
		}

		if err := tx.
			Where("role_id = ?", role.ID).
			Delete(&models.RolePermission{}).
			Error; err != nil {
			return err
		}

		if len(*permissionIDs) == 0 {
			return nil
		}

		rolePermissions := make([]models.RolePermission, 0, len(*permissionIDs))

		for _, permissionID := range *permissionIDs {
			rolePermissions = append(rolePermissions, models.RolePermission{
				RoleID:       role.ID,
				PermissionID: permissionID,
			})
		}

		return tx.Create(&rolePermissions).Error
	})
}

func (r *RoleRepository) Delete(role *models.Role) error {
	return r.db.Delete(role).Error
}

func (r *RoleRepository) FindPermissionsByRoleID(roleID int64) ([]models.Permission, error) {
	var permissions []models.Permission

	err := r.db.
		Table("permissions").
		Select("permissions.id, permissions.code, permissions.description").
		Joins("JOIN role_permissions ON role_permissions.permission_id = permissions.id").
		Where("role_permissions.role_id = ?", roleID).
		Order("permissions.code ASC").
		Find(&permissions).
		Error

	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (r *RoleRepository) ExistsByName(name string, excludeID *int64) (bool, error) {
	var count int64

	query := r.db.
		Model(&models.Role{}).
		Where("name = ?", name)

	if excludeID != nil {
		query = query.Where("id <> ?", *excludeID)
	}

	if err := query.Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *RoleRepository) PermissionsExist(permissionIDs []int64) (bool, error) {
	if len(permissionIDs) == 0 {
		return true, nil
	}

	var count int64

	err := r.db.
		Model(&models.Permission{}).
		Where("id IN ?", permissionIDs).
		Distinct("id").
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count == int64(len(permissionIDs)), nil
}

func (r *RoleRepository) HasUsers(roleID int64) (bool, error) {
	var count int64

	err := r.db.
		Table("user_roles").
		Where("role_id = ?", roleID).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
