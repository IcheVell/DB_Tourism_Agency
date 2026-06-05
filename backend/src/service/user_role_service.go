package service

import (
	"errors"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"

	"gorm.io/gorm"
)

var (
	ErrUserRoleNotFound      = errors.New("user role not found")
	ErrUserRoleAlreadyExists = errors.New("user role already exists")
)

type UserRoleService struct {
	userRoleRepository *repository.UserRoleRepository
}

func NewUserRoleService(userRoleRepository *repository.UserRoleRepository) *UserRoleService {
	return &UserRoleService{
		userRoleRepository: userRoleRepository,
	}
}

func (s *UserRoleService) Create(request dto.CreateUserRoleRequest) (*dto.UserRoleResponse, error) {
	if request.UserID <= 0 || request.RoleID <= 0 {
		return nil, ErrInvalidInput
	}

	userExists, err := s.userRoleRepository.UserExists(request.UserID)
	if err != nil {
		return nil, err
	}

	if !userExists {
		return nil, ErrUserNotFound
	}

	roleExists, err := s.userRoleRepository.RoleExists(request.RoleID)
	if err != nil {
		return nil, err
	}

	if !roleExists {
		return nil, ErrRoleNotFound
	}

	userHasRole, err := s.userRoleRepository.UserHasRole(request.UserID)
	if err != nil {
		return nil, err
	}

	if userHasRole {
		return nil, ErrUserRoleAlreadyExists
	}

	userRole := &models.UserRole{
		UserID: request.UserID,
		RoleID: request.RoleID,
	}

	if err := s.userRoleRepository.Create(userRole); err != nil {
		return nil, err
	}

	createdUserRole, err := s.userRoleRepository.FindByIDs(request.UserID, request.RoleID)
	if err != nil {
		return nil, err
	}

	return userRoleToResponse(createdUserRole), nil
}

func (s *UserRoleService) FindByIDs(userID int64, roleID int64) (*dto.UserRoleResponse, error) {
	if userID <= 0 || roleID <= 0 {
		return nil, ErrInvalidInput
	}

	userRole, err := s.userRoleRepository.FindByIDs(userID, roleID)
	if err != nil {
		return nil, err
	}

	if userRole == nil {
		return nil, ErrUserRoleNotFound
	}

	return userRoleToResponse(userRole), nil
}

func (s *UserRoleService) FindByUserID(userID int64) (*dto.UserRoleResponse, error) {
	if userID <= 0 {
		return nil, ErrInvalidInput
	}

	userRole, err := s.userRoleRepository.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	if userRole == nil {
		return nil, ErrUserRoleNotFound
	}

	return userRoleToResponse(userRole), nil
}

func (s *UserRoleService) FindAll(page int, limit int, userID *int64) (*dto.UserRoleListResponse, error) {
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

	var userRoles []models.UserRole
	var total int64
	var err error

	if userID != nil {
		if *userID <= 0 {
			return nil, ErrInvalidInput
		}

		userExists, err := s.userRoleRepository.UserExists(*userID)
		if err != nil {
			return nil, err
		}

		if !userExists {
			return nil, ErrUserNotFound
		}

		userRoles, total, err = s.userRoleRepository.FindAllByUserID(*userID, limit, offset)
	} else {
		userRoles, total, err = s.userRoleRepository.FindAll(limit, offset)
	}

	if err != nil {
		return nil, err
	}

	items := make([]dto.UserRoleResponse, 0, len(userRoles))

	for i := range userRoles {
		items = append(items, *userRoleToResponse(&userRoles[i]))
	}

	return &dto.UserRoleListResponse{
		Items: items,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *UserRoleService) Update(
	oldUserID int64,
	oldRoleID int64,
	request dto.UpdateUserRoleRequest,
) (*dto.UserRoleResponse, error) {
	if oldUserID <= 0 || oldRoleID <= 0 {
		return nil, ErrInvalidInput
	}

	oldUserRole, err := s.userRoleRepository.FindByIDs(oldUserID, oldRoleID)
	if err != nil {
		return nil, err
	}

	if oldUserRole == nil {
		return nil, ErrUserRoleNotFound
	}

	newUserID := oldUserID
	newRoleID := oldRoleID

	if request.UserID != nil {
		if *request.UserID <= 0 {
			return nil, ErrInvalidInput
		}

		newUserID = *request.UserID
	}

	if request.RoleID != nil {
		if *request.RoleID <= 0 {
			return nil, ErrInvalidInput
		}

		newRoleID = *request.RoleID
	}

	userExists, err := s.userRoleRepository.UserExists(newUserID)
	if err != nil {
		return nil, err
	}

	if !userExists {
		return nil, ErrUserNotFound
	}

	roleExists, err := s.userRoleRepository.RoleExists(newRoleID)
	if err != nil {
		return nil, err
	}

	if !roleExists {
		return nil, ErrRoleNotFound
	}

	if newUserID != oldUserID {
		userHasRole, err := s.userRoleRepository.UserHasRole(newUserID)
		if err != nil {
			return nil, err
		}

		if userHasRole {
			return nil, ErrUserRoleAlreadyExists
		}
	}

	if newUserID != oldUserID || newRoleID != oldRoleID {
		exists, err := s.userRoleRepository.ExistsByIDs(newUserID, newRoleID)
		if err != nil {
			return nil, err
		}

		if exists {
			return nil, ErrUserRoleAlreadyExists
		}
	}

	newUserRole := &models.UserRole{
		UserID: newUserID,
		RoleID: newRoleID,
	}

	if err := s.userRoleRepository.Update(oldUserID, oldRoleID, newUserRole); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserRoleNotFound
		}

		return nil, err
	}

	updatedUserRole, err := s.userRoleRepository.FindByIDs(newUserID, newRoleID)
	if err != nil {
		return nil, err
	}

	return userRoleToResponse(updatedUserRole), nil
}

func (s *UserRoleService) Delete(userID int64, roleID int64) error {
	if userID <= 0 || roleID <= 0 {
		return ErrInvalidInput
	}

	userRole, err := s.userRoleRepository.FindByIDs(userID, roleID)
	if err != nil {
		return err
	}

	if userRole == nil {
		return ErrUserRoleNotFound
	}

	return s.userRoleRepository.Delete(userRole)
}

func userRoleToResponse(userRole *models.UserRole) *dto.UserRoleResponse {
	response := &dto.UserRoleResponse{
		UserID: userRole.UserID,
		RoleID: userRole.RoleID,
	}

	if userRole.User != nil {
		response.User = &dto.UserBriefResponse{
			ID:    userRole.User.ID,
			Login: userRole.User.Login,
			Email: userRole.User.Email,
		}
	}

	if userRole.Role != nil {
		response.Role = &dto.RoleBriefResponse{
			ID:          userRole.Role.ID,
			Name:        userRole.Role.Name,
			Description: userRole.Role.Description,
		}
	}

	return response
}
