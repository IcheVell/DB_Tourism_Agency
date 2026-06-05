package service

import (
	"errors"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"

	"gorm.io/gorm"
)

var (
	ErrRolePermissionNotFound      = errors.New("role permission not found")
	ErrRolePermissionAlreadyExists = errors.New("role permission already exists")
)

type RolePermissionService struct {
	rolePermissionRepository *repository.RolePermissionRepository
}

func NewRolePermissionService(rolePermissionRepository *repository.RolePermissionRepository) *RolePermissionService {
	return &RolePermissionService{
		rolePermissionRepository: rolePermissionRepository,
	}
}

func (s *RolePermissionService) Create(request dto.CreateRolePermissionRequest) (*dto.RolePermissionResponse, error) {
	if request.RoleID <= 0 || request.PermissionID <= 0 {
		return nil, ErrInvalidInput
	}

	roleExists, err := s.rolePermissionRepository.RoleExists(request.RoleID)
	if err != nil {
		return nil, err
	}

	if !roleExists {
		return nil, ErrRoleNotFound
	}

	permissionExists, err := s.rolePermissionRepository.PermissionExists(request.PermissionID)
	if err != nil {
		return nil, err
	}

	if !permissionExists {
		return nil, ErrPermissionNotFound
	}

	exists, err := s.rolePermissionRepository.ExistsByIDs(request.RoleID, request.PermissionID)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrRolePermissionAlreadyExists
	}

	rolePermission := &models.RolePermission{
		RoleID:       request.RoleID,
		PermissionID: request.PermissionID,
	}

	if err := s.rolePermissionRepository.Create(rolePermission); err != nil {
		return nil, err
	}

	createdRolePermission, err := s.rolePermissionRepository.FindByIDs(request.RoleID, request.PermissionID)
	if err != nil {
		return nil, err
	}

	return rolePermissionToResponse(createdRolePermission), nil
}

func (s *RolePermissionService) FindByIDs(roleID int64, permissionID int64) (*dto.RolePermissionResponse, error) {
	if roleID <= 0 || permissionID <= 0 {
		return nil, ErrInvalidInput
	}

	rolePermission, err := s.rolePermissionRepository.FindByIDs(roleID, permissionID)
	if err != nil {
		return nil, err
	}

	if rolePermission == nil {
		return nil, ErrRolePermissionNotFound
	}

	return rolePermissionToResponse(rolePermission), nil
}

func (s *RolePermissionService) FindAll(page int, limit int, roleID *int64) (*dto.RolePermissionListResponse, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 20
	}

	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit

	var rolePermissions []models.RolePermission
	var total int64
	var err error

	if roleID != nil {
		if *roleID <= 0 {
			return nil, ErrInvalidInput
		}

		roleExists, err := s.rolePermissionRepository.RoleExists(*roleID)
		if err != nil {
			return nil, err
		}

		if !roleExists {
			return nil, ErrRoleNotFound
		}

		rolePermissions, total, err = s.rolePermissionRepository.FindByRoleID(*roleID, limit, offset)
	} else {
		rolePermissions, total, err = s.rolePermissionRepository.FindAll(limit, offset)
	}

	if err != nil {
		return nil, err
	}

	items := make([]dto.RolePermissionResponse, 0, len(rolePermissions))

	for i := range rolePermissions {
		items = append(items, *rolePermissionToResponse(&rolePermissions[i]))
	}

	return &dto.RolePermissionListResponse{
		Items: items,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *RolePermissionService) Update(
	oldRoleID int64,
	oldPermissionID int64,
	request dto.UpdateRolePermissionRequest,
) (*dto.RolePermissionResponse, error) {
	if oldRoleID <= 0 || oldPermissionID <= 0 {
		return nil, ErrInvalidInput
	}

	oldRolePermission, err := s.rolePermissionRepository.FindByIDs(oldRoleID, oldPermissionID)
	if err != nil {
		return nil, err
	}

	if oldRolePermission == nil {
		return nil, ErrRolePermissionNotFound
	}

	newRoleID := oldRoleID
	newPermissionID := oldPermissionID

	if request.RoleID != nil {
		if *request.RoleID <= 0 {
			return nil, ErrInvalidInput
		}

		newRoleID = *request.RoleID
	}

	if request.PermissionID != nil {
		if *request.PermissionID <= 0 {
			return nil, ErrInvalidInput
		}

		newPermissionID = *request.PermissionID
	}

	roleExists, err := s.rolePermissionRepository.RoleExists(newRoleID)
	if err != nil {
		return nil, err
	}

	if !roleExists {
		return nil, ErrRoleNotFound
	}

	permissionExists, err := s.rolePermissionRepository.PermissionExists(newPermissionID)
	if err != nil {
		return nil, err
	}

	if !permissionExists {
		return nil, ErrPermissionNotFound
	}

	if newRoleID != oldRoleID || newPermissionID != oldPermissionID {
		exists, err := s.rolePermissionRepository.ExistsByIDs(newRoleID, newPermissionID)
		if err != nil {
			return nil, err
		}

		if exists {
			return nil, ErrRolePermissionAlreadyExists
		}
	}

	newRolePermission := &models.RolePermission{
		RoleID:       newRoleID,
		PermissionID: newPermissionID,
	}

	if err := s.rolePermissionRepository.Update(oldRoleID, oldPermissionID, newRolePermission); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRolePermissionNotFound
		}

		return nil, err
	}

	updatedRolePermission, err := s.rolePermissionRepository.FindByIDs(newRoleID, newPermissionID)
	if err != nil {
		return nil, err
	}

	return rolePermissionToResponse(updatedRolePermission), nil
}

func (s *RolePermissionService) Delete(roleID int64, permissionID int64) error {
	if roleID <= 0 || permissionID <= 0 {
		return ErrInvalidInput
	}

	rolePermission, err := s.rolePermissionRepository.FindByIDs(roleID, permissionID)
	if err != nil {
		return err
	}

	if rolePermission == nil {
		return ErrRolePermissionNotFound
	}

	return s.rolePermissionRepository.Delete(rolePermission)
}

func rolePermissionToResponse(rolePermission *models.RolePermission) *dto.RolePermissionResponse {
	response := &dto.RolePermissionResponse{
		RoleID:       rolePermission.RoleID,
		PermissionID: rolePermission.PermissionID,
	}

	if rolePermission.Role != nil {
		response.Role = &dto.RoleBriefResponse{
			ID:          rolePermission.Role.ID,
			Name:        rolePermission.Role.Name,
			Description: rolePermission.Role.Description,
		}
	}

	if rolePermission.Permission != nil {
		response.Permission = &dto.PermissionResponse{
			ID:          rolePermission.Permission.ID,
			Code:        rolePermission.Permission.Code,
			Description: rolePermission.Permission.Description,
		}
	}

	return response
}
