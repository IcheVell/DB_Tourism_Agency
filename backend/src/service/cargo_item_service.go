package service

import (
	"errors"
	"strings"
	"time"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"
)

var (
	ErrCargoItemNotFound                  = errors.New("cargo item not found")
	ErrCargoItemInvalidItemNumber         = errors.New("item_number is required")
	ErrCargoItemInvalidWeightKg           = errors.New("weight_kg must be positive")
	ErrCargoItemInvalidVolumetricWeightKg = errors.New("volumetric_weight_kg must be non-negative")
	ErrCargoItemInvalidPlacesCount        = errors.New("places_count must be positive")
	ErrCargoItemInvalidPackagedAt         = errors.New("packaged_at must be RFC3339 datetime")
	ErrCargoItemInvalidCargoTypeID        = errors.New("cargo_type_id must be positive")
	ErrCargoItemInvalidCargoStatementID   = errors.New("cargo_statement_id must be positive")
)

type CargoItemService struct {
	repo *repository.CargoItemRepository
}

func NewCargoItemService(repo *repository.CargoItemRepository) *CargoItemService {
	return &CargoItemService{
		repo: repo,
	}
}

func (s *CargoItemService) Create(req dto.CreateCargoItemRequest) (*dto.CargoItemResponse, error) {
	item, err := buildCargoItem(
		req.ItemNumber,
		req.WeightKg,
		req.VolumetricWeightKg,
		req.PlacesCount,
		req.Marking,
		req.PackagedAt,
		req.CargoTypeID,
		req.CargoStatementID,
	)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(item); err != nil {
		return nil, err
	}

	return toCargoItemResponse(item), nil
}

func (s *CargoItemService) GetByID(id int64) (*dto.CargoItemResponse, error) {
	item, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, ErrCargoItemNotFound
	}

	return toCargoItemResponse(item), nil
}

func (s *CargoItemService) List(page int, pageSize int) (*dto.CargoItemListResponse, error) {
	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 20
	}

	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	items, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	responseItems := make([]dto.CargoItemResponse, 0, len(items))
	for _, item := range items {
		responseItems = append(responseItems, *toCargoItemResponse(&item))
	}

	return &dto.CargoItemListResponse{
		Items:    responseItems,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *CargoItemService) Update(id int64, req dto.UpdateCargoItemRequest) (*dto.CargoItemResponse, error) {
	item, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, ErrCargoItemNotFound
	}

	updatedItem, err := buildCargoItem(
		req.ItemNumber,
		req.WeightKg,
		req.VolumetricWeightKg,
		req.PlacesCount,
		req.Marking,
		req.PackagedAt,
		req.CargoTypeID,
		req.CargoStatementID,
	)
	if err != nil {
		return nil, err
	}

	item.ItemNumber = updatedItem.ItemNumber
	item.WeightKg = updatedItem.WeightKg
	item.VolumetricWeightKg = updatedItem.VolumetricWeightKg
	item.PlacesCount = updatedItem.PlacesCount
	item.Marking = updatedItem.Marking
	item.PackagedAt = updatedItem.PackagedAt
	item.CargoTypeID = updatedItem.CargoTypeID
	item.CargoStatementID = updatedItem.CargoStatementID

	if err := s.repo.Update(item); err != nil {
		return nil, err
	}

	return toCargoItemResponse(item), nil
}

func (s *CargoItemService) Delete(id int64) error {
	item, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if item == nil {
		return ErrCargoItemNotFound
	}

	return s.repo.Delete(item)
}

func buildCargoItem(
	itemNumber string,
	weightKg float64,
	volumetricWeightKg float64,
	placesCount int,
	marking *string,
	packagedAtRaw *string,
	cargoTypeID int64,
	cargoStatementID int64,
) (*models.CargoItem, error) {
	normalizedItemNumber := strings.TrimSpace(itemNumber)
	if normalizedItemNumber == "" {
		return nil, ErrCargoItemInvalidItemNumber
	}

	if weightKg <= 0 {
		return nil, ErrCargoItemInvalidWeightKg
	}

	if volumetricWeightKg < 0 {
		return nil, ErrCargoItemInvalidVolumetricWeightKg
	}

	if placesCount <= 0 {
		return nil, ErrCargoItemInvalidPlacesCount
	}

	normalizedMarking := normalizeOptionalString(marking)

	packagedAt, err := parseOptionalDateTimeRFC3339(packagedAtRaw)
	if err != nil {
		return nil, ErrCargoItemInvalidPackagedAt
	}

	if cargoTypeID <= 0 {
		return nil, ErrCargoItemInvalidCargoTypeID
	}

	if cargoStatementID <= 0 {
		return nil, ErrCargoItemInvalidCargoStatementID
	}

	return &models.CargoItem{
		ItemNumber:         normalizedItemNumber,
		WeightKg:           weightKg,
		VolumetricWeightKg: volumetricWeightKg,
		PlacesCount:        placesCount,
		Marking:            normalizedMarking,
		PackagedAt:         packagedAt,
		CargoTypeID:        cargoTypeID,
		CargoStatementID:   cargoStatementID,
	}, nil
}

func toCargoItemResponse(item *models.CargoItem) *dto.CargoItemResponse {
	return &dto.CargoItemResponse{
		ID:                 item.ID,
		ItemNumber:         item.ItemNumber,
		WeightKg:           item.WeightKg,
		VolumetricWeightKg: item.VolumetricWeightKg,
		PlacesCount:        item.PlacesCount,
		Marking:            item.Marking,
		PackagedAt:         formatOptionalCargoItemDateTime(item.PackagedAt),
		CargoTypeID:        item.CargoTypeID,
		CargoStatementID:   item.CargoStatementID,
	}
}

func formatOptionalCargoItemDateTime(value *time.Time) *string {
	if value == nil {
		return nil
	}

	formatted := value.Format(time.RFC3339)
	return &formatted
}
