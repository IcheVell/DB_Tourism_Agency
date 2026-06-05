package service

import (
	"errors"
	"time"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"
)

var (
	ErrFinancialOperationNotFound          = errors.New("financial operation not found")
	ErrFinancialOperationInvalidAmount     = errors.New("financial operation amount must be positive")
	ErrFinancialOperationInvalidCategoryID = errors.New("financial category id must be positive")
	ErrFinancialOperationInvalidSource     = errors.New("financial operation must have exactly one source")
)

type FinancialOperationService struct {
	repo *repository.FinancialOperationRepository
}

func NewFinancialOperationService(repo *repository.FinancialOperationRepository) *FinancialOperationService {
	return &FinancialOperationService{
		repo: repo,
	}
}

func (s *FinancialOperationService) Create(req dto.CreateFinancialOperationRequest) (*dto.FinancialOperationResponse, error) {
	if req.Amount <= 0 {
		return nil, ErrFinancialOperationInvalidAmount
	}

	if req.FinancialCategoryID <= 0 {
		return nil, ErrFinancialOperationInvalidCategoryID
	}

	if countFinancialOperationSources(
		req.FlightID,
		req.VisaID,
		req.ExcursionScheduleID,
		req.ExcursionBookingID,
		req.CargoShipmentID,
		req.CargoStatementID,
		req.AccommodationID,
	) != 1 {
		return nil, ErrFinancialOperationInvalidSource
	}

	operationAt := time.Now().UTC()
	if req.OperationAt != nil {
		operationAt = *req.OperationAt
	}

	operation := &models.FinancialOperation{
		Amount:              req.Amount,
		OperationAt:         operationAt,
		Description:         normalizeOptionalString(req.Description),
		FinancialCategoryID: req.FinancialCategoryID,

		FlightID:            req.FlightID,
		VisaID:              req.VisaID,
		ExcursionScheduleID: req.ExcursionScheduleID,
		ExcursionBookingID:  req.ExcursionBookingID,
		CargoShipmentID:     req.CargoShipmentID,
		CargoStatementID:    req.CargoStatementID,
		AccommodationID:     req.AccommodationID,
	}

	if err := s.repo.Create(operation); err != nil {
		return nil, err
	}

	return toFinancialOperationResponse(operation), nil
}

func (s *FinancialOperationService) GetByID(id int64) (*dto.FinancialOperationResponse, error) {
	operation, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if operation == nil {
		return nil, ErrFinancialOperationNotFound
	}

	return toFinancialOperationResponse(operation), nil
}

func (s *FinancialOperationService) List(page int, pageSize int) (*dto.FinancialOperationListResponse, error) {
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

	operations, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.FinancialOperationResponse, 0, len(operations))
	for _, operation := range operations {
		items = append(items, *toFinancialOperationResponse(&operation))
	}

	return &dto.FinancialOperationListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *FinancialOperationService) Update(id int64, req dto.UpdateFinancialOperationRequest) (*dto.FinancialOperationResponse, error) {
	if req.Amount <= 0 {
		return nil, ErrFinancialOperationInvalidAmount
	}

	if req.FinancialCategoryID <= 0 {
		return nil, ErrFinancialOperationInvalidCategoryID
	}

	if countFinancialOperationSources(
		req.FlightID,
		req.VisaID,
		req.ExcursionScheduleID,
		req.ExcursionBookingID,
		req.CargoShipmentID,
		req.CargoStatementID,
		req.AccommodationID,
	) != 1 {
		return nil, ErrFinancialOperationInvalidSource
	}

	operation, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if operation == nil {
		return nil, ErrFinancialOperationNotFound
	}

	operation.Amount = req.Amount
	operation.Description = normalizeOptionalString(req.Description)
	operation.FinancialCategoryID = req.FinancialCategoryID

	if req.OperationAt != nil {
		operation.OperationAt = *req.OperationAt
	}

	operation.FlightID = req.FlightID
	operation.VisaID = req.VisaID
	operation.ExcursionScheduleID = req.ExcursionScheduleID
	operation.ExcursionBookingID = req.ExcursionBookingID
	operation.CargoShipmentID = req.CargoShipmentID
	operation.CargoStatementID = req.CargoStatementID
	operation.AccommodationID = req.AccommodationID

	if err := s.repo.Update(operation); err != nil {
		return nil, err
	}

	return toFinancialOperationResponse(operation), nil
}

func (s *FinancialOperationService) Delete(id int64) error {
	operation, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if operation == nil {
		return ErrFinancialOperationNotFound
	}

	return s.repo.Delete(operation)
}

func toFinancialOperationResponse(operation *models.FinancialOperation) *dto.FinancialOperationResponse {
	return &dto.FinancialOperationResponse{
		ID:                  operation.ID,
		Amount:              operation.Amount,
		OperationAt:         operation.OperationAt,
		Description:         operation.Description,
		FinancialCategoryID: operation.FinancialCategoryID,

		FlightID:            operation.FlightID,
		VisaID:              operation.VisaID,
		ExcursionScheduleID: operation.ExcursionScheduleID,
		ExcursionBookingID:  operation.ExcursionBookingID,
		CargoShipmentID:     operation.CargoShipmentID,
		CargoStatementID:    operation.CargoStatementID,
		AccommodationID:     operation.AccommodationID,
	}
}

func countFinancialOperationSources(sourceIDs ...*int64) int {
	count := 0

	for _, sourceID := range sourceIDs {
		if sourceID != nil && *sourceID > 0 {
			count++
		}
	}

	return count
}
