package repository

import (
	"errors"
	"time"

	"TouristAgencyApp/src/dto"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

type AuthUserRecord struct {
	ID           int64
	Login        string
	Email        string
	PasswordHash string
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) FindUserByLogin(login string) (*AuthUserRecord, error) {
	var user AuthUserRecord

	err := r.db.
		Table("users").
		Select("id, login, email, password_hash").
		Where("login = ?", login).
		First(&user).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) ExistsByLogin(login string) (bool, error) {
	var count int64

	err := r.db.
		Table("users").
		Where("login = ?", login).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *AuthRepository) ExistsByEmail(email string) (bool, error) {
	var count int64

	err := r.db.
		Table("users").
		Where("email = ?", email).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *AuthRepository) RegisterTouristUser(
	login string,
	email string,
	passwordHash string,
	firstName string,
	lastName string,
	middleName *string,
	sex string,
	birthDate time.Time,
) (*dto.RegisterResponse, error) {
	var response dto.RegisterResponse

	err := r.db.Transaction(func(tx *gorm.DB) error {
		var touristRoleID int64

		if err := tx.
			Table("roles").
			Select("id").
			Where("name = ?", "tourist").
			Scan(&touristRoleID).
			Error; err != nil {
			return err
		}

		if touristRoleID == 0 {
			return gorm.ErrRecordNotFound
		}

		var userID int64

		if err := tx.Raw(`
			INSERT INTO users (
				login,
				email,
				password_hash
			)
			VALUES (?, ?, ?)
			RETURNING id
		`, login, email, passwordHash).Scan(&userID).Error; err != nil {
			return err
		}

		var touristID int64

		if err := tx.Raw(`
			INSERT INTO tourists (
				first_name,
				last_name,
				middle_name,
				sex,
				birth_date,
				user_id
			)
			VALUES (?, ?, ?, ?, ?, ?)
			RETURNING id
		`, firstName, lastName, middleName, sex, birthDate, userID).Scan(&touristID).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
			INSERT INTO user_roles (
				user_id,
				role_id
			)
			VALUES (?, ?)
		`, userID, touristRoleID).Error; err != nil {
			return err
		}

		response = dto.RegisterResponse{
			ID:        userID,
			Login:     login,
			Email:     email,
			Role:      "tourist",
			TouristID: touristID,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *AuthRepository) FindMeByUserID(userID int64) (*dto.MeResponse, error) {
	var response dto.MeResponse

	err := r.db.
		Table("users").
		Select("id, login, email").
		Where("id = ?", userID).
		First(&response).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var roles []string

	if err := r.db.
		Table("roles").
		Select("roles.name").
		Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Order("roles.name ASC").
		Pluck("roles.name", &roles).
		Error; err != nil {
		return nil, err
	}

	var permissions []string

	if err := r.db.
		Table("permissions").
		Select("DISTINCT permissions.code").
		Joins("JOIN role_permissions ON role_permissions.permission_id = permissions.id").
		Joins("JOIN user_roles ON user_roles.role_id = role_permissions.role_id").
		Where("user_roles.user_id = ?", userID).
		Order("permissions.code ASC").
		Pluck("permissions.code", &permissions).
		Error; err != nil {
		return nil, err
	}

	response.Roles = roles
	response.Permissions = permissions

	return &response, nil
}
