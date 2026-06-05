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
	ErrExcursionBookingNotFound                 = errors.New("excursion booking not found")
	ErrExcursionBookingInvalidBookedAt          = errors.New("booked_at must be RFC3339 datetime")
	ErrExcursionBookingInvalidStatus            = errors.New("status must be booked, visited or cancelled")
	ErrExcursionBookingInvalidTouristRating     = errors.New("tourist_rating must be between 1 and 5")
	ErrExcursionBookingRatingRequiredForVisited = errors.New("visited booking requires tourist_rating")
	ErrExcursionBookingRatingOnlyForVisited     = errors.New("tourist_rating is allowed only for visited status")
	ErrExcursionBookingInvalidScheduleID        = errors.New("excursion_schedule_id must be positive")
	ErrExcursionBookingInvalidGroupMemberID     = errors.New("group_member_id must be positive")
)

type ExcursionBookingService struct {
	repo *repository.ExcursionBookingRepository
}

func NewExcursionBookingService(repo *repository.ExcursionBookingRepository) *ExcursionBookingService {
	return &ExcursionBookingService{
		repo: repo,
	}
}

func (s *ExcursionBookingService) Create(req dto.CreateExcursionBookingRequest) (*dto.ExcursionBookingResponse, error) {
	booking, err := buildExcursionBooking(
		req.BookedAt,
		req.TouristRating,
		req.Status,
		req.ExcursionScheduleID,
		req.GroupMemberID,
	)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(booking); err != nil {
		return nil, err
	}

	return toExcursionBookingResponse(booking), nil
}

func (s *ExcursionBookingService) GetByID(id int64) (*dto.ExcursionBookingResponse, error) {
	booking, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if booking == nil {
		return nil, ErrExcursionBookingNotFound
	}

	return toExcursionBookingResponse(booking), nil
}

func (s *ExcursionBookingService) List(page int, pageSize int) (*dto.ExcursionBookingListResponse, error) {
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

	bookings, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.ExcursionBookingResponse, 0, len(bookings))
	for _, booking := range bookings {
		items = append(items, *toExcursionBookingResponse(&booking))
	}

	return &dto.ExcursionBookingListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *ExcursionBookingService) Update(id int64, req dto.UpdateExcursionBookingRequest) (*dto.ExcursionBookingResponse, error) {
	booking, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if booking == nil {
		return nil, ErrExcursionBookingNotFound
	}

	updatedBooking, err := buildExcursionBooking(
		req.BookedAt,
		req.TouristRating,
		req.Status,
		req.ExcursionScheduleID,
		req.GroupMemberID,
	)
	if err != nil {
		return nil, err
	}

	booking.BookedAt = updatedBooking.BookedAt
	booking.TouristRating = updatedBooking.TouristRating
	booking.Status = updatedBooking.Status
	booking.ExcursionScheduleID = updatedBooking.ExcursionScheduleID
	booking.GroupMemberID = updatedBooking.GroupMemberID

	if err := s.repo.Update(booking); err != nil {
		return nil, err
	}

	return toExcursionBookingResponse(booking), nil
}

func (s *ExcursionBookingService) Delete(id int64) error {
	booking, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if booking == nil {
		return ErrExcursionBookingNotFound
	}

	return s.repo.Delete(booking)
}

func buildExcursionBooking(
	bookedAtRaw string,
	touristRating *int,
	status string,
	excursionScheduleID int64,
	groupMemberID int64,
) (*models.ExcursionBooking, error) {
	bookedAt, err := parseDateTimeRFC3339(bookedAtRaw)
	if err != nil {
		return nil, ErrExcursionBookingInvalidBookedAt
	}

	normalizedStatus := strings.TrimSpace(status)
	if !isValidExcursionBookingStatus(normalizedStatus) {
		return nil, ErrExcursionBookingInvalidStatus
	}

	if touristRating != nil && (*touristRating < 1 || *touristRating > 5) {
		return nil, ErrExcursionBookingInvalidTouristRating
	}

	if normalizedStatus == "visited" && touristRating == nil {
		return nil, ErrExcursionBookingRatingRequiredForVisited
	}

	if normalizedStatus != "visited" && touristRating != nil {
		return nil, ErrExcursionBookingRatingOnlyForVisited
	}

	if excursionScheduleID <= 0 {
		return nil, ErrExcursionBookingInvalidScheduleID
	}

	if groupMemberID <= 0 {
		return nil, ErrExcursionBookingInvalidGroupMemberID
	}

	return &models.ExcursionBooking{
		BookedAt:            bookedAt,
		TouristRating:       touristRating,
		Status:              normalizedStatus,
		ExcursionScheduleID: excursionScheduleID,
		GroupMemberID:       groupMemberID,
	}, nil
}

func toExcursionBookingResponse(booking *models.ExcursionBooking) *dto.ExcursionBookingResponse {
	return &dto.ExcursionBookingResponse{
		ID:                  booking.ID,
		BookedAt:            booking.BookedAt.Format(time.RFC3339),
		TouristRating:       booking.TouristRating,
		Status:              booking.Status,
		ExcursionScheduleID: booking.ExcursionScheduleID,
		GroupMemberID:       booking.GroupMemberID,
	}
}

func isValidExcursionBookingStatus(status string) bool {
	switch status {
	case "booked", "visited", "cancelled":
		return true
	default:
		return false
	}
}
