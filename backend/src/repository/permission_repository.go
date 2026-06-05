package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type PermissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{
		db: db,
	}
}

func (r *PermissionRepository) Create(permission *models.Permission) error {
	return r.db.Create(permission).Error
}

func (r *PermissionRepository) FindByID(id int64) (*models.Permission, error) {
	var permission models.Permission

	err := r.db.
		Where("id = ?", id).
		First(&permission).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &permission, nil
}

func (r *PermissionRepository) FindAll(limit int, offset int) ([]models.Permission, int64, error) {
	var permissions []models.Permission
	var total int64

	if err := r.db.Model(&models.Permission{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("code ASC").
		Limit(limit).
		Offset(offset).
		Find(&permissions).
		Error

	if err != nil {
		return nil, 0, err
	}

	return permissions, total, nil
}

func (r *PermissionRepository) Update(permission *models.Permission) error {
	return r.db.Save(permission).Error
}

func (r *PermissionRepository) Delete(permission *models.Permission) error {
	return r.db.Delete(permission).Error
}

func (r *PermissionRepository) ExistsByCode(code string, excludeID *int64) (bool, error) {
	var count int64

	query := r.db.
		Model(&models.Permission{}).
		Where("code = ?", code)

	if excludeID != nil {
		query = query.Where("id <> ?", *excludeID)
	}

	if err := query.Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *PermissionRepository) HasRoles(permissionID int64) (bool, error) {
	var count int64

	err := r.db.
		Table("role_permissions").
		Where("permission_id = ?", permissionID).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
