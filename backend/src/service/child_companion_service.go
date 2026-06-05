package service

import (
	"errors"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"
)

var (
	ErrChildCompanionNotFound         = errors.New("child companion not found")
	ErrChildCompanionInvalidAdultID   = errors.New("adult_group_member_id must be positive")
	ErrChildCompanionInvalidChildID   = errors.New("child_group_member_id must be positive")
	ErrChildCompanionAdultEqualsChild = errors.New("adult_group_member_id and child_group_member_id must be different")
)

type ChildCompanionService struct {
	repo *repository.ChildCompanionRepository
}

func NewChildCompanionService(repo *repository.ChildCompanionRepository) *ChildCompanionService {
	return &ChildCompanionService{
		repo: repo,
	}
}

func (s *ChildCompanionService) Create(req dto.CreateChildCompanionRequest) (*dto.ChildCompanionResponse, error) {
	companion, err := buildChildCompanion(req.AdultGroupMemberID, req.ChildGroupMemberID)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(companion); err != nil {
		return nil, err
	}

	return toChildCompanionResponse(companion), nil
}

func (s *ChildCompanionService) GetByIDs(
	AdultGroupMemberID int64,
	childGroupMemberID int64,
) (*dto.ChildCompanionResponse, error) {
	companion, err := s.repo.FindByIDs(AdultGroupMemberID, childGroupMemberID)
	if err != nil {
		return nil, err
	}

	if companion == nil {
		return nil, ErrChildCompanionNotFound
	}

	return toChildCompanionResponse(companion), nil
}

func (s *ChildCompanionService) List(page int, pageSize int) (*dto.ChildCompanionListResponse, error) {
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

	companions, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.ChildCompanionResponse, 0, len(companions))
	for _, companion := range companions {
		items = append(items, *toChildCompanionResponse(&companion))
	}

	return &dto.ChildCompanionListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *ChildCompanionService) Update(
	oldParentGroupMemberID int64,
	oldChildGroupMemberID int64,
	req dto.UpdateChildCompanionRequest,
) (*dto.ChildCompanionResponse, error) {
	oldCompanion, err := s.repo.FindByIDs(oldParentGroupMemberID, oldChildGroupMemberID)
	if err != nil {
		return nil, err
	}

	if oldCompanion == nil {
		return nil, ErrChildCompanionNotFound
	}

	newCompanion, err := buildChildCompanion(req.AdultGroupMemberID, req.ChildGroupMemberID)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Replace(oldCompanion, newCompanion); err != nil {
		return nil, err
	}

	return toChildCompanionResponse(newCompanion), nil
}

func (s *ChildCompanionService) Delete(AdultGroupMemberID int64, childGroupMemberID int64) error {
	companion, err := s.repo.FindByIDs(AdultGroupMemberID, childGroupMemberID)
	if err != nil {
		return err
	}

	if companion == nil {
		return ErrChildCompanionNotFound
	}

	return s.repo.Delete(companion)
}

func buildChildCompanion(adultGroupMemberID int64, childGroupMemberID int64) (*models.ChildCompanion, error) {
	if adultGroupMemberID <= 0 {
		return nil, ErrChildCompanionInvalidAdultID
	}

	if childGroupMemberID <= 0 {
		return nil, ErrChildCompanionInvalidChildID
	}

	if adultGroupMemberID == childGroupMemberID {
		return nil, ErrChildCompanionAdultEqualsChild
	}

	return &models.ChildCompanion{
		AdultGroupMemberID: adultGroupMemberID,
		ChildGroupMemberID: childGroupMemberID,
	}, nil
}

func toChildCompanionResponse(companion *models.ChildCompanion) *dto.ChildCompanionResponse {
	return &dto.ChildCompanionResponse{
		AdultGroupMemberID: companion.AdultGroupMemberID,
		ChildGroupMemberID: companion.ChildGroupMemberID,
	}
}
