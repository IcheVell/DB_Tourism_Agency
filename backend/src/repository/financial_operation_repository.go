package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type FinancialOperationRepository struct {
	db *gorm.DB
}

func NewFinancialOperationRepository(db *gorm.DB) *FinancialOperationRepository {
	return &FinancialOperationRepository{
		db: db,
	}
}

func (r *FinancialOperationRepository) Create(operation *models.FinancialOperation) error {
	return r.db.Create(operation).Error
}

func (r *FinancialOperationRepository) FindByID(id int64) (*models.FinancialOperation, error) {
	var operation models.FinancialOperation

	err := r.db.
		Where("id = ?", id).
		First(&operation).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &operation, nil
}

func (r *FinancialOperationRepository) FindAll(limit int, offset int) ([]models.FinancialOperation, int64, error) {
	var operations []models.FinancialOperation
	var total int64

	if err := r.db.Model(&models.FinancialOperation{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&operations).
		Error

	if err != nil {
		return nil, 0, err
	}

	return operations, total, nil
}

func (r *FinancialOperationRepository) Update(operation *models.FinancialOperation) error {
	return r.db.Save(operation).Error
}

func (r *FinancialOperationRepository) Delete(operation *models.FinancialOperation) error {
	return r.db.Delete(operation).Error
}
