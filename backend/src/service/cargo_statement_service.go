package service

import (
	"errors"
	"strings"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"
)

var (
	ErrCargoStatementNotFound             = errors.New("cargo statement not found")
	ErrCargoStatementInvalidStatus        = errors.New("cargo statement status is invalid")
	ErrCargoStatementInvalidGroupMemberID = errors.New("group_member_id must be positive")
)

type CargoStatementService struct {
	repo *repository.CargoStatementRepository
}

func NewCargoStatementService(repo *repository.CargoStatementRepository) *CargoStatementService {
	return &CargoStatementService{
		repo: repo,
	}
}

func (s *CargoStatementService) Create(req dto.CreateCargoStatementRequest) (*dto.CargoStatementResponse, error) {
	statement, err := buildCargoStatement(req.Status, req.GroupMemberID)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(statement); err != nil {
		return nil, err
	}

	return toCargoStatementResponse(statement), nil
}

func (s *CargoStatementService) GetByID(id int64) (*dto.CargoStatementResponse, error) {
	statement, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if statement == nil {
		return nil, ErrCargoStatementNotFound
	}

	return toCargoStatementResponse(statement), nil
}

func (s *CargoStatementService) List(page int, pageSize int) (*dto.CargoStatementListResponse, error) {
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

	statements, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.CargoStatementResponse, 0, len(statements))
	for _, statement := range statements {
		items = append(items, *toCargoStatementResponse(&statement))
	}

	return &dto.CargoStatementListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *CargoStatementService) Update(id int64, req dto.UpdateCargoStatementRequest) (*dto.CargoStatementResponse, error) {
	statement, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if statement == nil {
		return nil, ErrCargoStatementNotFound
	}

	updatedStatement, err := buildCargoStatement(req.Status, req.GroupMemberID)
	if err != nil {
		return nil, err
	}

	statement.Status = updatedStatement.Status
	statement.GroupMemberID = updatedStatement.GroupMemberID

	if err := s.repo.Update(statement); err != nil {
		return nil, err
	}

	return toCargoStatementResponse(statement), nil
}

func (s *CargoStatementService) Delete(id int64) error {
	statement, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if statement == nil {
		return ErrCargoStatementNotFound
	}

	return s.repo.Delete(statement)
}

func buildCargoStatement(status string, groupMemberID int64) (*models.CargoStatement, error) {
	normalizedStatus := strings.TrimSpace(status)

	if !isValidCargoStatementStatus(normalizedStatus) {
		return nil, ErrCargoStatementInvalidStatus
	}

	if groupMemberID <= 0 {
		return nil, ErrCargoStatementInvalidGroupMemberID
	}

	return &models.CargoStatement{
		Status:        normalizedStatus,
		GroupMemberID: groupMemberID,
	}, nil
}

func toCargoStatementResponse(statement *models.CargoStatement) *dto.CargoStatementResponse {
	return &dto.CargoStatementResponse{
		ID:            statement.ID,
		Status:        statement.Status,
		GroupMemberID: statement.GroupMemberID,
	}
}

func isValidCargoStatementStatus(status string) bool {
	switch status {
	case "draft", "weighed", "packed", "ready_for_shipment", "shipped", "cancelled":
		return true
	default:
		return false
	}
}
