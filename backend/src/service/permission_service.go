package service

import (
	"errors"
	"strings"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"
)

var (
	ErrPermissionAlreadyExists = errors.New("permission already exists")
	ErrPermissionHasRoles      = errors.New("permission has roles")
)

type PermissionService struct {
	permissionRepository *repository.PermissionRepository
}

func NewPermissionService(permissionRepository *repository.PermissionRepository) *PermissionService {
	return &PermissionService{
		permissionRepository: permissionRepository,
	}
}

func (s *PermissionService) Create(request dto.CreatePermissionRequest) (*dto.PermissionResponse, error) {
	code := strings.TrimSpace(request.Code)

	if code == "" || len(code) > 255 {
		return nil, ErrInvalidInput
	}

	description := normalizePermissionStringPointer(request.Description)

	exists, err := s.permissionRepository.ExistsByCode(code, nil)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrPermissionAlreadyExists
	}

	permission := &models.Permission{
		Code:        code,
		Description: description,
	}

	if err := s.permissionRepository.Create(permission); err != nil {
		return nil, err
	}

	return permissionToResponse(permission), nil
}

func (s *PermissionService) FindByID(id int64) (*dto.PermissionResponse, error) {
	permission, err := s.permissionRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if permission == nil {
		return nil, ErrPermissionNotFound
	}

	return permissionToResponse(permission), nil
}

func (s *PermissionService) FindAll(page int, limit int) (*dto.PermissionListResponse, error) {
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

	permissions, total, err := s.permissionRepository.FindAll(limit, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.PermissionResponse, 0, len(permissions))

	for i := range permissions {
		items = append(items, *permissionToResponse(&permissions[i]))
	}

	return &dto.PermissionListResponse{
		Items: items,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *PermissionService) Update(id int64, request dto.UpdatePermissionRequest) (*dto.PermissionResponse, error) {
	permission, err := s.permissionRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if permission == nil {
		return nil, ErrPermissionNotFound
	}

	if request.Code != nil {
		code := strings.TrimSpace(*request.Code)

		if code == "" || len(code) > 255 {
			return nil, ErrInvalidInput
		}

		exists, err := s.permissionRepository.ExistsByCode(code, &id)
		if err != nil {
			return nil, err
		}

		if exists {
			return nil, ErrPermissionAlreadyExists
		}

		permission.Code = code
	}

	if request.Description != nil {
		permission.Description = normalizePermissionStringPointer(request.Description)
	}

	if err := s.permissionRepository.Update(permission); err != nil {
		return nil, err
	}

	return permissionToResponse(permission), nil
}

func (s *PermissionService) Delete(id int64) error {
	permission, err := s.permissionRepository.FindByID(id)
	if err != nil {
		return err
	}

	if permission == nil {
		return ErrPermissionNotFound
	}

	hasRoles, err := s.permissionRepository.HasRoles(id)
	if err != nil {
		return err
	}

	if hasRoles {
		return ErrPermissionHasRoles
	}

	return s.permissionRepository.Delete(permission)
}

func permissionToResponse(permission *models.Permission) *dto.PermissionResponse {
	return &dto.PermissionResponse{
		ID:          permission.ID,
		Code:        permission.Code,
		Description: permission.Description,
	}
}

func normalizePermissionStringPointer(value *string) *string {
	if value == nil {
		return nil
	}

	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		return nil
	}

	return &trimmed
}
