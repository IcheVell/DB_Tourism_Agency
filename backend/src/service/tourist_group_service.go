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
	ErrTouristGroupNotFound             = errors.New("tourist group not found")
	ErrTouristGroupInvalidName          = errors.New("tourist group name is required")
	ErrTouristGroupInvalidArrivalDate   = errors.New("arrival_date must be RFC3339 datetime")
	ErrTouristGroupInvalidDepartureDate = errors.New("departure_date must be RFC3339 datetime")
	ErrTouristGroupInvalidDateRange     = errors.New("arrival_date must be before departure_date")
)

type TouristGroupService struct {
	repo *repository.TouristGroupRepository
}

func NewTouristGroupService(repo *repository.TouristGroupRepository) *TouristGroupService {
	return &TouristGroupService{
		repo: repo,
	}
}

func (s *TouristGroupService) Create(req dto.CreateTouristGroupRequest) (*dto.TouristGroupResponse, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, ErrTouristGroupInvalidName
	}

	arrivalDate, err := parseDateTimeRFC3339(req.ArrivalDate)
	if err != nil {
		return nil, ErrTouristGroupInvalidArrivalDate
	}

	departureDate, err := parseDateTimeRFC3339(req.DepartureDate)
	if err != nil {
		return nil, ErrTouristGroupInvalidDepartureDate
	}

	if !arrivalDate.Before(departureDate) {
		return nil, ErrTouristGroupInvalidDateRange
	}

	group := &models.TouristGroup{
		Name:          name,
		ArrivalDate:   arrivalDate,
		DepartureDate: departureDate,
	}

	if err := s.repo.Create(group); err != nil {
		return nil, err
	}

	return toTouristGroupResponse(group), nil
}

func (s *TouristGroupService) GetByID(id int64) (*dto.TouristGroupResponse, error) {
	group, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if group == nil {
		return nil, ErrTouristGroupNotFound
	}

	return toTouristGroupResponse(group), nil
}

func (s *TouristGroupService) List(page int, pageSize int) (*dto.TouristGroupListResponse, error) {
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

	groups, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.TouristGroupResponse, 0, len(groups))
	for _, group := range groups {
		items = append(items, *toTouristGroupResponse(&group))
	}

	return &dto.TouristGroupListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *TouristGroupService) Update(id int64, req dto.UpdateTouristGroupRequest) (*dto.TouristGroupResponse, error) {
	group, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if group == nil {
		return nil, ErrTouristGroupNotFound
	}

	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, ErrTouristGroupInvalidName
	}

	arrivalDate, err := parseDateTimeRFC3339(req.ArrivalDate)
	if err != nil {
		return nil, ErrTouristGroupInvalidArrivalDate
	}

	departureDate, err := parseDateTimeRFC3339(req.DepartureDate)
	if err != nil {
		return nil, ErrTouristGroupInvalidDepartureDate
	}

	if !arrivalDate.Before(departureDate) {
		return nil, ErrTouristGroupInvalidDateRange
	}

	group.Name = name
	group.ArrivalDate = arrivalDate
	group.DepartureDate = departureDate

	if err := s.repo.Update(group); err != nil {
		return nil, err
	}

	return toTouristGroupResponse(group), nil
}

func (s *TouristGroupService) Delete(id int64) error {
	group, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if group == nil {
		return ErrTouristGroupNotFound
	}

	return s.repo.Delete(group)
}

func toTouristGroupResponse(group *models.TouristGroup) *dto.TouristGroupResponse {
	return &dto.TouristGroupResponse{
		ID:            group.ID,
		Name:          group.Name,
		ArrivalDate:   group.ArrivalDate.Format(time.RFC3339),
		DepartureDate: group.DepartureDate.Format(time.RFC3339),
	}
}

func parseDateTimeRFC3339(value string) (time.Time, error) {
	return time.Parse(time.RFC3339, strings.TrimSpace(value))
}
