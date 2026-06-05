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
	ErrCargoShipmentNotFound                = errors.New("cargo shipment not found")
	ErrCargoShipmentInvalidStatus           = errors.New("status is required")
	ErrCargoShipmentInvalidShippedAt        = errors.New("shipped_at must be RFC3339 datetime")
	ErrCargoShipmentInvalidCargoStatementID = errors.New("cargo_statement_id must be positive")
	ErrCargoShipmentInvalidFlightID         = errors.New("flight_id must be positive")
)

type CargoShipmentService struct {
	repo *repository.CargoShipmentRepository
}

func NewCargoShipmentService(repo *repository.CargoShipmentRepository) *CargoShipmentService {
	return &CargoShipmentService{
		repo: repo,
	}
}

func (s *CargoShipmentService) Create(req dto.CreateCargoShipmentRequest) (*dto.CargoShipmentResponse, error) {
	shipment, err := buildCargoShipment(
		req.ShippedAt,
		req.Status,
		req.CargoStatementID,
		req.FlightID,
	)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(shipment); err != nil {
		return nil, err
	}

	return toCargoShipmentResponse(shipment), nil
}

func (s *CargoShipmentService) GetByID(id int64) (*dto.CargoShipmentResponse, error) {
	shipment, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if shipment == nil {
		return nil, ErrCargoShipmentNotFound
	}

	return toCargoShipmentResponse(shipment), nil
}

func (s *CargoShipmentService) List(page int, pageSize int) (*dto.CargoShipmentListResponse, error) {
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

	shipments, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.CargoShipmentResponse, 0, len(shipments))
	for _, shipment := range shipments {
		items = append(items, *toCargoShipmentResponse(&shipment))
	}

	return &dto.CargoShipmentListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *CargoShipmentService) Update(id int64, req dto.UpdateCargoShipmentRequest) (*dto.CargoShipmentResponse, error) {
	shipment, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if shipment == nil {
		return nil, ErrCargoShipmentNotFound
	}

	updatedShipment, err := buildCargoShipment(
		req.ShippedAt,
		req.Status,
		req.CargoStatementID,
		req.FlightID,
	)
	if err != nil {
		return nil, err
	}

	shipment.ShippedAt = updatedShipment.ShippedAt
	shipment.Status = updatedShipment.Status
	shipment.CargoStatementID = updatedShipment.CargoStatementID
	shipment.FlightID = updatedShipment.FlightID

	if err := s.repo.Update(shipment); err != nil {
		return nil, err
	}

	return toCargoShipmentResponse(shipment), nil
}

func (s *CargoShipmentService) Delete(id int64) error {
	shipment, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if shipment == nil {
		return ErrCargoShipmentNotFound
	}

	return s.repo.Delete(shipment)
}

func buildCargoShipment(
	shippedAtRaw *string,
	status string,
	cargoStatementID int64,
	flightID int64,
) (*models.CargoShipment, error) {
	normalizedStatus := strings.TrimSpace(status)

	if normalizedStatus == "" {
		return nil, ErrCargoShipmentInvalidStatus
	}

	shippedAt, err := parseOptionalDateTimeRFC3339(shippedAtRaw)
	if err != nil {
		return nil, ErrCargoShipmentInvalidShippedAt
	}

	if cargoStatementID <= 0 {
		return nil, ErrCargoShipmentInvalidCargoStatementID
	}

	if flightID <= 0 {
		return nil, ErrCargoShipmentInvalidFlightID
	}

	return &models.CargoShipment{
		ShippedAt:        shippedAt,
		Status:           normalizedStatus,
		CargoStatementID: cargoStatementID,
		FlightID:         flightID,
	}, nil
}

func toCargoShipmentResponse(shipment *models.CargoShipment) *dto.CargoShipmentResponse {
	return &dto.CargoShipmentResponse{
		ID:               shipment.ID,
		ShippedAt:        formatOptionalCargoShipmentDateTime(shipment.ShippedAt),
		Status:           shipment.Status,
		CargoStatementID: shipment.CargoStatementID,
		FlightID:         shipment.FlightID,
	}
}

func formatOptionalCargoShipmentDateTime(value *time.Time) *string {
	if value == nil {
		return nil
	}

	formatted := value.Format(time.RFC3339)
	return &formatted
}
