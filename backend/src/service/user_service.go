package service

import (
	"context"
	"errors"
	"net/mail"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/repository"
)

var (
	ErrLoginAlreadyUsed = errors.New("login already used")
	ErrEmailAlreadyUsed = errors.New("email already used")
	ErrUserNotFound     = errors.New("user not found")
)

var loginRegexp = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{2,49}$`)

type UserService interface {
	Create(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error)
	GetByID(ctx context.Context, id uint64) (*dto.UserResponse, error)
	List(ctx context.Context, page int, limit int) (*dto.UserListResponse, error)
	Update(ctx context.Context, id uint64, req dto.UpdateUserRequest) (*dto.UserResponse, error)
	Delete(ctx context.Context, id uint64) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) Create(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error) {
	req.Login = strings.TrimSpace(req.Login)
	req.Email = strings.TrimSpace(req.Email)

	if err := validateLogin(req.Login); err != nil {
		return nil, err
	}

	if err := validateEmail(req.Email); err != nil {
		return nil, err
	}

	if err := validatePassword(req.Password); err != nil {
		return nil, err
	}

	if req.RoleID == 0 {
		return nil, ErrInvalidInput
	}

	roleExists, err := s.userRepository.RoleExists(ctx, req.RoleID)
	if err != nil {
		return nil, err
	}

	if !roleExists {
		return nil, ErrRoleNotFound
	}

	loginUsed, err := s.userRepository.ExistsByLogin(ctx, req.Login, nil)
	if err != nil {
		return nil, err
	}

	if loginUsed {
		return nil, ErrLoginAlreadyUsed
	}

	emailUsed, err := s.userRepository.ExistsByEmail(ctx, req.Email, nil)
	if err != nil {
		return nil, err
	}

	if emailUsed {
		return nil, ErrEmailAlreadyUsed
	}

	passwordHash, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &repository.UserRecord{
		Login:        req.Login,
		Email:        req.Email,
		PasswordHash: passwordHash,
	}

	if err := s.userRepository.Create(ctx, user, req.RoleID); err != nil {
		return nil, err
	}

	createdUser, err := s.userRepository.GetByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return toUserResponse(createdUser), nil
}

func (s *userService) GetByID(ctx context.Context, id uint64) (*dto.UserResponse, error) {
	user, err := s.userRepository.GetByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return toUserResponse(user), nil
}

func (s *userService) List(ctx context.Context, page int, limit int) (*dto.UserListResponse, error) {
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

	users, total, err := s.userRepository.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.UserResponse, 0, len(users))
	for i := range users {
		items = append(items, *toUserResponse(&users[i]))
	}

	return &dto.UserListResponse{
		Items: items,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *userService) Update(ctx context.Context, id uint64, req dto.UpdateUserRequest) (*dto.UserResponse, error) {
	currentUser, err := s.userRepository.GetByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	if req.Login != nil {
		login := strings.TrimSpace(*req.Login)

		if err := validateLogin(login); err != nil {
			return nil, err
		}

		if login != currentUser.Login {
			loginUsed, err := s.userRepository.ExistsByLogin(ctx, login, &id)
			if err != nil {
				return nil, err
			}

			if loginUsed {
				return nil, ErrLoginAlreadyUsed
			}
		}

		currentUser.Login = login
	}

	if req.Email != nil {
		email := strings.TrimSpace(*req.Email)

		if err := validateEmail(email); err != nil {
			return nil, err
		}

		if email != currentUser.Email {
			emailUsed, err := s.userRepository.ExistsByEmail(ctx, email, &id)
			if err != nil {
				return nil, err
			}

			if emailUsed {
				return nil, ErrEmailAlreadyUsed
			}
		}

		currentUser.Email = email
	}

	if req.Password != nil {
		if err := validatePassword(*req.Password); err != nil {
			return nil, err
		}

		passwordHash, err := hashPassword(*req.Password)
		if err != nil {
			return nil, err
		}

		currentUser.PasswordHash = passwordHash
	} else {
		currentUser.PasswordHash = ""
	}

	if req.RoleID != nil {
		if *req.RoleID == 0 {
			return nil, ErrInvalidInput
		}

		roleExists, err := s.userRepository.RoleExists(ctx, *req.RoleID)
		if err != nil {
			return nil, err
		}

		if !roleExists {
			return nil, ErrRoleNotFound
		}
	}

	if err := s.userRepository.Update(ctx, currentUser, req.RoleID); err != nil {
		return nil, err
	}

	updatedUser, err := s.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return toUserResponse(updatedUser), nil
}

func (s *userService) Delete(ctx context.Context, id uint64) error {
	err := s.userRepository.Delete(ctx, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrUserNotFound
	}

	return err
}

func validateLogin(login string) error {
	if !loginRegexp.MatchString(login) {
		return ErrInvalidInput
	}

	return nil
}

func validateEmail(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return ErrInvalidInput
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return ErrInvalidInput
	}

	return nil
}

func hashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

func toUserResponse(user *repository.UserRecord) *dto.UserResponse {
	response := &dto.UserResponse{
		ID:        user.ID,
		Login:     user.Login,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	if user.Role != nil {
		response.Role = &dto.RoleResponse{
			ID:          int64(user.Role.ID),
			Name:        user.Role.Name,
			Description: user.Role.Description,
		}
	}

	return response
}
