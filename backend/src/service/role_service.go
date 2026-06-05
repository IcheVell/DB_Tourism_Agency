package service

import (
	"errors"
	"strings"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"
)

var (
	ErrRoleNotFound       = errors.New("role not found")
	ErrRoleAlreadyExists  = errors.New("role already exists")
	ErrRoleHasUsers       = errors.New("role has users")
	ErrPermissionNotFound = errors.New("permission not found")
)

type RoleService struct {
	roleRepository *repository.RoleRepository
}

func NewRoleService(roleRepository *repository.RoleRepository) *RoleService {
	return &RoleService{
		roleRepository: roleRepository,
	}
}

func (s *RoleService) Create(request dto.CreateRoleRequest) (*dto.RoleResponse, error) {
	name := strings.TrimSpace(request.Name)

	if name == "" || len(name) > 255 {
		return nil, ErrInvalidInput
	}

	description := normalizeRoleStringPointer(request.Description)
	permissionIDs := uniqueRoleInt64(request.PermissionIDs)

	exists, err := s.roleRepository.ExistsByName(name, nil)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrRoleAlreadyExists
	}

	permissionsExist, err := s.roleRepository.PermissionsExist(permissionIDs)
	if err != nil {
		return nil, err
	}

	if !permissionsExist {
		return nil, ErrPermissionNotFound
	}

	role := &models.Role{
		Name:        name,
		Description: description,
	}

	if err := s.roleRepository.Create(role, permissionIDs); err != nil {
		return nil, err
	}

	return s.buildRoleResponse(role)
}

func (s *RoleService) FindByID(id int64) (*dto.RoleResponse, error) {
	role, err := s.roleRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if role == nil {
		return nil, ErrRoleNotFound
	}

	return s.buildRoleResponse(role)
}

func (s *RoleService) FindAll(page int, limit int) (*dto.RoleListResponse, error) {
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

	roles, total, err := s.roleRepository.FindAll(limit, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.RoleResponse, 0, len(roles))

	for i := range roles {
		response, err := s.buildRoleResponse(&roles[i])
		if err != nil {
			return nil, err
		}

		items = append(items, *response)
	}

	return &dto.RoleListResponse{
		Items: items,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *RoleService) Update(id int64, request dto.UpdateRoleRequest) (*dto.RoleResponse, error) {
	role, err := s.roleRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if role == nil {
		return nil, ErrRoleNotFound
	}

	if request.Name != nil {
		name := strings.TrimSpace(*request.Name)

		if name == "" || len(name) > 255 {
			return nil, ErrInvalidInput
		}

		exists, err := s.roleRepository.ExistsByName(name, &id)
		if err != nil {
			return nil, err
		}

		if exists {
			return nil, ErrRoleAlreadyExists
		}

		role.Name = name
	}

	if request.Description != nil {
		role.Description = normalizeRoleStringPointer(request.Description)
	}

	var permissionIDs *[]int64

	if request.PermissionIDs != nil {
		ids := uniqueRoleInt64(*request.PermissionIDs)

		permissionsExist, err := s.roleRepository.PermissionsExist(ids)
		if err != nil {
			return nil, err
		}

		if !permissionsExist {
			return nil, ErrPermissionNotFound
		}

		permissionIDs = &ids
	}

	if err := s.roleRepository.Update(role, permissionIDs); err != nil {
		return nil, err
	}

	updatedRole, err := s.roleRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s.buildRoleResponse(updatedRole)
}

func (s *RoleService) Delete(id int64) error {
	role, err := s.roleRepository.FindByID(id)
	if err != nil {
		return err
	}

	if role == nil {
		return ErrRoleNotFound
	}

	hasUsers, err := s.roleRepository.HasUsers(id)
	if err != nil {
		return err
	}

	if hasUsers {
		return ErrRoleHasUsers
	}

	return s.roleRepository.Delete(role)
}

func (s *RoleService) buildRoleResponse(role *models.Role) (*dto.RoleResponse, error) {
	permissions, err := s.roleRepository.FindPermissionsByRoleID(role.ID)
	if err != nil {
		return nil, err
	}

	permissionResponses := make([]dto.PermissionResponse, 0, len(permissions))

	for i := range permissions {
		permissionResponses = append(permissionResponses, dto.PermissionResponse{
			ID:          permissions[i].ID,
			Code:        permissions[i].Code,
			Description: permissions[i].Description,
		})
	}

	return &dto.RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		Permissions: permissionResponses,
	}, nil
}

func normalizeRoleStringPointer(value *string) *string {
	if value == nil {
		return nil
	}

	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		return nil
	}

	return &trimmed
}

func uniqueRoleInt64(values []int64) []int64 {
	seen := make(map[int64]struct{}, len(values))
	result := make([]int64, 0, len(values))

	for _, value := range values {
		if value <= 0 {
			continue
		}

		if _, exists := seen[value]; exists {
			continue
		}

		seen[value] = struct{}{}
		result = append(result, value)
	}

	return result
}
