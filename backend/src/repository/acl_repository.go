package repository

import "gorm.io/gorm"

type ACLRepository struct {
	db *gorm.DB
}

func NewACLRepository(db *gorm.DB) *ACLRepository {
	return &ACLRepository{
		db: db,
	}
}

func (r *ACLRepository) HasPermission(userID int64, permissionCode string) (bool, error) {
	var count int64

	err := r.db.
		Table("permissions").
		Joins("JOIN role_permissions ON role_permissions.permission_id = permissions.id").
		Joins("JOIN user_roles ON user_roles.role_id = role_permissions.role_id").
		Where("user_roles.user_id = ?", userID).
		Where("permissions.code = ?", permissionCode).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
