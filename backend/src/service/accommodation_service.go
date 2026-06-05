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
	ErrAccommodationNotFound             = errors.New("accommodation not found")
	ErrAccommodationInvalidStatus        = errors.New("accommodation status is invalid")
	ErrAccommodationInvalidCheckInAt     = errors.New("check_in_at must be RFC3339 datetime")
	ErrAccommodationInvalidCheckOutAt    = errors.New("check_out_at must be RFC3339 datetime")
	ErrAccommodationInvalidDateRange     = errors.New("check_out_at must be after check_in_at")
	ErrAccommodationCheckOutRequired     = errors.New("checked_out accommodation requires check_out_at")
	ErrAccommodationInvalidGroupMemberID = errors.New("group_member_id must be positive")
	ErrAccommodationInvalidHotelRoomID   = errors.New("hotel_room_id must be positive")
)

type AccommodationService struct {
	repo *repository.AccommodationRepository
}

func NewAccommodationService(repo *repository.AccommodationRepository) *AccommodationService {
	return &AccommodationService{
		repo: repo,
	}
}

func (s *AccommodationService) Create(req dto.CreateAccommodationRequest) (*dto.AccommodationResponse, error) {
	accommodation, err := buildAccommodationFromCreateRequest(req)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(accommodation); err != nil {
		return nil, err
	}

	return toAccommodationResponse(accommodation), nil
}

func (s *AccommodationService) GetByID(id int64) (*dto.AccommodationResponse, error) {
	accommodation, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if accommodation == nil {
		return nil, ErrAccommodationNotFound
	}

	return toAccommodationResponse(accommodation), nil
}

func (s *AccommodationService) List(page int, pageSize int) (*dto.AccommodationListResponse, error) {
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

	accommodations, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.AccommodationResponse, 0, len(accommodations))
	for _, accommodation := range accommodations {
		items = append(items, *toAccommodationResponse(&accommodation))
	}

	return &dto.AccommodationListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *AccommodationService) Update(id int64, req dto.UpdateAccommodationRequest) (*dto.AccommodationResponse, error) {
	accommodation, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if accommodation == nil {
		return nil, ErrAccommodationNotFound
	}

	updatedAccommodation, err := buildAccommodationFromUpdateRequest(req)
	if err != nil {
		return nil, err
	}

	accommodation.Status = updatedAccommodation.Status
	accommodation.CheckInAt = updatedAccommodation.CheckInAt
	accommodation.CheckOutAt = updatedAccommodation.CheckOutAt
	accommodation.GroupMemberID = updatedAccommodation.GroupMemberID
	accommodation.HotelRoomID = updatedAccommodation.HotelRoomID

	if err := s.repo.Update(accommodation); err != nil {
		return nil, err
	}

	return toAccommodationResponse(accommodation), nil
}

func (s *AccommodationService) Delete(id int64) error {
	accommodation, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if accommodation == nil {
		return ErrAccommodationNotFound
	}

	return s.repo.Delete(accommodation)
}

func buildAccommodationFromCreateRequest(req dto.CreateAccommodationRequest) (*models.Accommodation, error) {
	return buildAccommodation(
		req.Status,
		req.CheckInAt,
		req.CheckOutAt,
		req.GroupMemberID,
		req.HotelRoomID,
	)
}

func buildAccommodationFromUpdateRequest(req dto.UpdateAccommodationRequest) (*models.Accommodation, error) {
	return buildAccommodation(
		req.Status,
		req.CheckInAt,
		req.CheckOutAt,
		req.GroupMemberID,
		req.HotelRoomID,
	)
}

func buildAccommodation(
	status string,
	checkInAtRaw string,
	checkOutAtRaw *string,
	groupMemberID int64,
	hotelRoomID int64,
) (*models.Accommodation, error) {
	normalizedStatus := strings.TrimSpace(status)

	if !isValidAccommodationStatus(normalizedStatus) {
		return nil, ErrAccommodationInvalidStatus
	}

	checkInAt, err := parseDateTimeRFC3339(checkInAtRaw)
	if err != nil {
		return nil, ErrAccommodationInvalidCheckInAt
	}

	checkOutAt, err := parseOptionalDateTimeRFC3339(checkOutAtRaw)
	if err != nil {
		return nil, ErrAccommodationInvalidCheckOutAt
	}

	if checkOutAt != nil && !checkOutAt.After(checkInAt) {
		return nil, ErrAccommodationInvalidDateRange
	}

	if normalizedStatus == "checked_out" && checkOutAt == nil {
		return nil, ErrAccommodationCheckOutRequired
	}

	if groupMemberID <= 0 {
		return nil, ErrAccommodationInvalidGroupMemberID
	}

	if hotelRoomID <= 0 {
		return nil, ErrAccommodationInvalidHotelRoomID
	}

	return &models.Accommodation{
		Status:        normalizedStatus,
		CheckInAt:     checkInAt,
		CheckOutAt:    checkOutAt,
		GroupMemberID: groupMemberID,
		HotelRoomID:   hotelRoomID,
	}, nil
}

func toAccommodationResponse(accommodation *models.Accommodation) *dto.AccommodationResponse {
	return &dto.AccommodationResponse{
		ID:            accommodation.ID,
		Status:        accommodation.Status,
		CheckInAt:     accommodation.CheckInAt.Format(time.RFC3339),
		CheckOutAt:    formatOptionalDateTimeRFC3339(accommodation.CheckOutAt),
		GroupMemberID: accommodation.GroupMemberID,
		HotelRoomID:   accommodation.HotelRoomID,
	}
}

func isValidAccommodationStatus(status string) bool {
	switch status {
	case "reserved", "checked_in", "checked_out", "cancelled":
		return true
	default:
		return false
	}
}
