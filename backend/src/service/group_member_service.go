package service

import (
	"errors"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"
)

var (
	ErrGroupMemberNotFound                 = errors.New("group member not found")
	ErrGroupMemberInvalidTouristGroupID    = errors.New("tourist_group_id must be positive")
	ErrGroupMemberInvalidTouristCategoryID = errors.New("tourist_category_id must be positive")
	ErrGroupMemberInvalidTouristID         = errors.New("tourist_id must be positive")
	ErrGroupMemberInvalidDesiredHotelID    = errors.New("desired_hotel_id must be positive")
)

type GroupMemberService struct {
	repo *repository.GroupMemberRepository
}

func NewGroupMemberService(repo *repository.GroupMemberRepository) *GroupMemberService {
	return &GroupMemberService{
		repo: repo,
	}
}

func (s *GroupMemberService) Create(req dto.CreateGroupMemberRequest) (*dto.GroupMemberResponse, error) {
	if err := validateGroupMemberIDs(
		req.TouristGroupID,
		req.TouristCategoryID,
		req.TouristID,
		req.DesiredHotelID,
	); err != nil {
		return nil, err
	}

	groupMember := &models.GroupMember{
		TouristGroupID:    req.TouristGroupID,
		TouristCategoryID: req.TouristCategoryID,
		TouristID:         req.TouristID,
		DesiredHotelID:    req.DesiredHotelID,
	}

	if err := s.repo.Create(groupMember); err != nil {
		return nil, err
	}

	return toGroupMemberResponse(groupMember), nil
}

func (s *GroupMemberService) GetByID(id int64) (*dto.GroupMemberResponse, error) {
	groupMember, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if groupMember == nil {
		return nil, ErrGroupMemberNotFound
	}

	return toGroupMemberResponse(groupMember), nil
}

func (s *GroupMemberService) List(page int, pageSize int) (*dto.GroupMemberListResponse, error) {
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

	groupMembers, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.GroupMemberResponse, 0, len(groupMembers))
	for _, groupMember := range groupMembers {
		items = append(items, *toGroupMemberResponse(&groupMember))
	}

	return &dto.GroupMemberListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *GroupMemberService) Update(id int64, req dto.UpdateGroupMemberRequest) (*dto.GroupMemberResponse, error) {
	if err := validateGroupMemberIDs(
		req.TouristGroupID,
		req.TouristCategoryID,
		req.TouristID,
		req.DesiredHotelID,
	); err != nil {
		return nil, err
	}

	groupMember, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if groupMember == nil {
		return nil, ErrGroupMemberNotFound
	}

	groupMember.TouristGroupID = req.TouristGroupID
	groupMember.TouristCategoryID = req.TouristCategoryID
	groupMember.TouristID = req.TouristID
	groupMember.DesiredHotelID = req.DesiredHotelID

	if err := s.repo.Update(groupMember); err != nil {
		return nil, err
	}

	return toGroupMemberResponse(groupMember), nil
}

func (s *GroupMemberService) Delete(id int64) error {
	groupMember, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if groupMember == nil {
		return ErrGroupMemberNotFound
	}

	return s.repo.Delete(groupMember)
}

func validateGroupMemberIDs(
	touristGroupID int64,
	touristCategoryID int64,
	touristID int64,
	desiredHotelID *int64,
) error {
	if touristGroupID <= 0 {
		return ErrGroupMemberInvalidTouristGroupID
	}

	if touristCategoryID <= 0 {
		return ErrGroupMemberInvalidTouristCategoryID
	}

	if touristID <= 0 {
		return ErrGroupMemberInvalidTouristID
	}

	if desiredHotelID != nil && *desiredHotelID <= 0 {
		return ErrGroupMemberInvalidDesiredHotelID
	}

	return nil
}

func toGroupMemberResponse(groupMember *models.GroupMember) *dto.GroupMemberResponse {
	return &dto.GroupMemberResponse{
		ID:                groupMember.ID,
		TouristGroupID:    groupMember.TouristGroupID,
		TouristCategoryID: groupMember.TouristCategoryID,
		TouristID:         groupMember.TouristID,
		DesiredHotelID:    groupMember.DesiredHotelID,
	}
}
