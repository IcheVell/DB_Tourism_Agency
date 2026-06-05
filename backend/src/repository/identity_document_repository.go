package repository

import (
	"errors"

	"TouristAgencyApp/src/models"

	"gorm.io/gorm"
)

type IdentityDocumentRepository struct {
	db *gorm.DB
}

func NewIdentityDocumentRepository(db *gorm.DB) *IdentityDocumentRepository {
	return &IdentityDocumentRepository{
		db: db,
	}
}

func (r *IdentityDocumentRepository) Create(document *models.IdentityDocument) error {
	return r.db.Create(document).Error
}

func (r *IdentityDocumentRepository) FindByID(id int64) (*models.IdentityDocument, error) {
	var document models.IdentityDocument

	err := r.db.
		Where("id = ?", id).
		First(&document).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &document, nil
}

func (r *IdentityDocumentRepository) FindAll(limit int, offset int) ([]models.IdentityDocument, int64, error) {
	var documents []models.IdentityDocument
	var total int64

	if err := r.db.Model(&models.IdentityDocument{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&documents).
		Error

	if err != nil {
		return nil, 0, err
	}

	return documents, total, nil
}

func (r *IdentityDocumentRepository) Update(document *models.IdentityDocument) error {
	return r.db.Save(document).Error
}

func (r *IdentityDocumentRepository) Delete(document *models.IdentityDocument) error {
	return r.db.Delete(document).Error
}
