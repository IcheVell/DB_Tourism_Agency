package repository

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserRecord struct {
	ID           uint64
	Login        string
	PasswordHash string
	Email        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Role         *RoleRecord `gorm:"-"`
}

func (UserRecord) TableName() string {
	return "users"
}

type RoleRecord struct {
	ID          uint64
	Name        string
	Description *string
}

func (RoleRecord) TableName() string {
	return "roles"
}

type UserRoleRecord struct {
	UserID uint64 `gorm:"column:user_id"`
	RoleID uint64 `gorm:"column:role_id"`
}

func (UserRoleRecord) TableName() string {
	return "user_roles"
}

type UserRepository interface {
	Create(ctx context.Context, user *UserRecord, roleID uint64) error
	GetByID(ctx context.Context, id uint64) (*UserRecord, error)
	List(ctx context.Context, limit int, offset int) ([]UserRecord, int64, error)
	Update(ctx context.Context, user *UserRecord, roleID *uint64) error
	Delete(ctx context.Context, id uint64) error

	ExistsByLogin(ctx context.Context, login string, excludeUserID *uint64) (bool, error)
	ExistsByEmail(ctx context.Context, email string, excludeUserID *uint64) (bool, error)
	RoleExists(ctx context.Context, roleID uint64) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *UserRecord, roleID uint64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		userRole := UserRoleRecord{
			UserID: user.ID,
			RoleID: roleID,
		}

		if err := tx.Create(&userRole).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *userRepository) GetByID(ctx context.Context, id uint64) (*UserRecord, error) {
	var user UserRecord

	if err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&user).Error; err != nil {
		return nil, err
	}

	role, err := r.getUserRole(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	user.Role = role

	return &user, nil
}

func (r *userRepository) List(ctx context.Context, limit int, offset int) ([]UserRecord, int64, error) {
	var total int64
	var users []UserRecord

	if err := r.db.WithContext(ctx).
		Model(&UserRecord{}).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).
		Order("id DESC").
		Limit(limit).
		Offset(offset).
		Find(&users).Error; err != nil {
		return nil, 0, err
	}

	for i := range users {
		role, err := r.getUserRole(ctx, users[i].ID)
		if err != nil {
			return nil, 0, err
		}
		users[i].Role = role
	}

	return users, total, nil
}

func (r *userRepository) Update(ctx context.Context, user *UserRecord, roleID *uint64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		updateData := map[string]any{
			"login":         user.Login,
			"email":         user.Email,
			"password_hash": user.PasswordHash,
		}

		if user.PasswordHash == "" {
			delete(updateData, "password_hash")
		}

		if err := tx.Model(&UserRecord{}).
			Where("id = ?", user.ID).
			Updates(updateData).Error; err != nil {
			return err
		}

		if roleID != nil {
			if err := tx.Where("user_id = ?", user.ID).
				Delete(&UserRoleRecord{}).Error; err != nil {
				return err
			}

			userRole := UserRoleRecord{
				UserID: user.ID,
				RoleID: *roleID,
			}

			if err := tx.Create(&userRole).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *userRepository) Delete(ctx context.Context, id uint64) error {
	result := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&UserRecord{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *userRepository) ExistsByLogin(ctx context.Context, login string, excludeUserID *uint64) (bool, error) {
	query := r.db.WithContext(ctx).
		Model(&UserRecord{}).
		Where("login = ?", login)

	if excludeUserID != nil {
		query = query.Where("id <> ?", *excludeUserID)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *userRepository) ExistsByEmail(ctx context.Context, email string, excludeUserID *uint64) (bool, error) {
	query := r.db.WithContext(ctx).
		Model(&UserRecord{}).
		Where("email = ?", email)

	if excludeUserID != nil {
		query = query.Where("id <> ?", *excludeUserID)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *userRepository) RoleExists(ctx context.Context, roleID uint64) (bool, error) {
	var count int64

	if err := r.db.WithContext(ctx).
		Model(&RoleRecord{}).
		Where("id = ?", roleID).
		Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *userRepository) getUserRole(ctx context.Context, userID uint64) (*RoleRecord, error) {
	var role RoleRecord

	err := r.db.WithContext(ctx).
		Table("roles").
		Select("roles.id, roles.name, roles.description").
		Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		First(&role).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &role, nil
}
