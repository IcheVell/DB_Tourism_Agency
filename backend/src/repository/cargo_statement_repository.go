package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type CargoStatementRepository struct {
	db *gorm.DB
}

func NewCargoStatementRepository(db *gorm.DB) *CargoStatementRepository {
	return &CargoStatementRepository{
		db: db,
	}
}

func (r *CargoStatementRepository) Create(statement *models.CargoStatement) error {
	return r.db.Create(statement).Error
}

func (r *CargoStatementRepository) FindByID(id int64) (*models.CargoStatement, error) {
	var statement models.CargoStatement

	err := r.db.
		Where("id = ?", id).
		First(&statement).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &statement, nil
}

func (r *CargoStatementRepository) FindAll(limit int, offset int) ([]models.CargoStatement, int64, error) {
	var statements []models.CargoStatement
	var total int64

	if err := r.db.Model(&models.CargoStatement{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&statements).
		Error

	if err != nil {
		return nil, 0, err
	}

	return statements, total, nil
}

func (r *CargoStatementRepository) Update(statement *models.CargoStatement) error {
	return r.db.Save(statement).Error
}

func (r *CargoStatementRepository) Delete(statement *models.CargoStatement) error {
	return r.db.Delete(statement).Error
}
