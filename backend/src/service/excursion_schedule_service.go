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
	ErrExcursionScheduleNotFound           = errors.New("excursion schedule not found")
	ErrExcursionScheduleInvalidPrice       = errors.New("price must be positive")
	ErrExcursionScheduleInvalidStartTime   = errors.New("start_time must be RFC3339 datetime")
	ErrExcursionScheduleInvalidEndTime     = errors.New("end_time must be RFC3339 datetime")
	ErrExcursionScheduleInvalidTimeRange   = errors.New("start_time must be before end_time")
	ErrExcursionScheduleInvalidCapacity    = errors.New("capacity must be positive")
	ErrExcursionScheduleInvalidStatus      = errors.New("status must be planned, completed or cancelled")
	ErrExcursionScheduleInvalidAgencyID    = errors.New("excursion_agency_id must be positive")
	ErrExcursionScheduleInvalidExcursionID = errors.New("excursion_id must be positive")
)

type ExcursionScheduleService struct {
	repo *repository.ExcursionScheduleRepository
}

func NewExcursionScheduleService(repo *repository.ExcursionScheduleRepository) *ExcursionScheduleService {
	return &ExcursionScheduleService{
		repo: repo,
	}
}

func (s *ExcursionScheduleService) Create(req dto.CreateExcursionScheduleRequest) (*dto.ExcursionScheduleResponse, error) {
	schedule, err := buildExcursionSchedule(
		req.Price,
		req.StartTime,
		req.EndTime,
		req.Capacity,
		req.Status,
		req.ExcursionAgencyID,
		req.ExcursionID,
	)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(schedule); err != nil {
		return nil, err
	}

	return toExcursionScheduleResponse(schedule), nil
}

func (s *ExcursionScheduleService) GetByID(id int64) (*dto.ExcursionScheduleResponse, error) {
	schedule, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if schedule == nil {
		return nil, ErrExcursionScheduleNotFound
	}

	return toExcursionScheduleResponse(schedule), nil
}

func (s *ExcursionScheduleService) List(page int, pageSize int) (*dto.ExcursionScheduleListResponse, error) {
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

	schedules, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.ExcursionScheduleResponse, 0, len(schedules))
	for _, schedule := range schedules {
		items = append(items, *toExcursionScheduleResponse(&schedule))
	}

	return &dto.ExcursionScheduleListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *ExcursionScheduleService) Update(id int64, req dto.UpdateExcursionScheduleRequest) (*dto.ExcursionScheduleResponse, error) {
	schedule, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if schedule == nil {
		return nil, ErrExcursionScheduleNotFound
	}

	updatedSchedule, err := buildExcursionSchedule(
		req.Price,
		req.StartTime,
		req.EndTime,
		req.Capacity,
		req.Status,
		req.ExcursionAgencyID,
		req.ExcursionID,
	)
	if err != nil {
		return nil, err
	}

	schedule.Price = updatedSchedule.Price
	schedule.StartTime = updatedSchedule.StartTime
	schedule.EndTime = updatedSchedule.EndTime
	schedule.Capacity = updatedSchedule.Capacity
	schedule.Status = updatedSchedule.Status
	schedule.ExcursionAgencyID = updatedSchedule.ExcursionAgencyID
	schedule.ExcursionID = updatedSchedule.ExcursionID

	if err := s.repo.Update(schedule); err != nil {
		return nil, err
	}

	return toExcursionScheduleResponse(schedule), nil
}

func (s *ExcursionScheduleService) Delete(id int64) error {
	schedule, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if schedule == nil {
		return ErrExcursionScheduleNotFound
	}

	return s.repo.Delete(schedule)
}

func buildExcursionSchedule(
	price float64,
	startTimeRaw string,
	endTimeRaw string,
	capacity int,
	status string,
	excursionAgencyID int64,
	excursionID int64,
) (*models.ExcursionSchedule, error) {
	if price <= 0 {
		return nil, ErrExcursionScheduleInvalidPrice
	}

	startTime, err := parseDateTimeRFC3339(startTimeRaw)
	if err != nil {
		return nil, ErrExcursionScheduleInvalidStartTime
	}

	endTime, err := parseDateTimeRFC3339(endTimeRaw)
	if err != nil {
		return nil, ErrExcursionScheduleInvalidEndTime
	}

	if !startTime.Before(endTime) {
		return nil, ErrExcursionScheduleInvalidTimeRange
	}

	if capacity <= 0 {
		return nil, ErrExcursionScheduleInvalidCapacity
	}

	normalizedStatus := strings.TrimSpace(status)
	if !isValidExcursionScheduleStatus(normalizedStatus) {
		return nil, ErrExcursionScheduleInvalidStatus
	}

	if excursionAgencyID <= 0 {
		return nil, ErrExcursionScheduleInvalidAgencyID
	}

	if excursionID <= 0 {
		return nil, ErrExcursionScheduleInvalidExcursionID
	}

	return &models.ExcursionSchedule{
		Price:             price,
		StartTime:         startTime,
		EndTime:           endTime,
		Capacity:          capacity,
		Status:            normalizedStatus,
		ExcursionAgencyID: excursionAgencyID,
		ExcursionID:       excursionID,
	}, nil
}

func toExcursionScheduleResponse(schedule *models.ExcursionSchedule) *dto.ExcursionScheduleResponse {
	return &dto.ExcursionScheduleResponse{
		ID:                schedule.ID,
		Price:             schedule.Price,
		StartTime:         schedule.StartTime.Format(time.RFC3339),
		EndTime:           schedule.EndTime.Format(time.RFC3339),
		Capacity:          schedule.Capacity,
		Status:            schedule.Status,
		ExcursionAgencyID: schedule.ExcursionAgencyID,
		ExcursionID:       schedule.ExcursionID,
	}
}

func isValidExcursionScheduleStatus(status string) bool {
	switch status {
	case "planned", "completed", "cancelled":
		return true
	default:
		return false
	}
}
