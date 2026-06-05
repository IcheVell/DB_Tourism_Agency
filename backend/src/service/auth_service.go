package service

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"
	"time"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrAuthInvalidInput        = errors.New("invalid auth input")
	ErrAuthInvalidCredentials  = errors.New("invalid credentials")
	ErrAuthUserNotFound        = errors.New("auth user not found")
	ErrAuthLoginAlreadyUsed    = errors.New("login already used")
	ErrAuthEmailAlreadyUsed    = errors.New("email already used")
	ErrAuthTouristRoleNotFound = errors.New("tourist role not found")
)

var authLoginRegexp = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{2,49}$`)

type AuthService struct {
	authRepository      *repository.AuthRepository
	jwtSecret           string
	jwtAccessTTLMinutes int
}

func NewAuthService(
	authRepository *repository.AuthRepository,
	jwtSecret string,
	jwtAccessTTLMinutes int,
) *AuthService {
	return &AuthService{
		authRepository:      authRepository,
		jwtSecret:           jwtSecret,
		jwtAccessTTLMinutes: jwtAccessTTLMinutes,
	}
}

func (s *AuthService) Register(request dto.RegisterRequest) (*dto.RegisterResponse, error) {
	login := strings.TrimSpace(request.Login)
	email := strings.TrimSpace(request.Email)
	password := request.Password

	firstName := strings.TrimSpace(request.FirstName)
	lastName := strings.TrimSpace(request.LastName)
	middleName := normalizeAuthStringPointer(request.MiddleName)
	sex := strings.TrimSpace(request.Sex)
	birthDateRaw := strings.TrimSpace(request.BirthDate)

	if !authLoginRegexp.MatchString(login) {
		return nil, ErrAuthInvalidInput
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return nil, ErrAuthInvalidInput
	}

	if len(password) < 8 {
		return nil, ErrAuthInvalidInput
	}

	if firstName == "" || lastName == "" {
		return nil, ErrAuthInvalidInput
	}

	if sex != "male" && sex != "female" {
		return nil, ErrAuthInvalidInput
	}

	birthDate, err := time.Parse("2006-01-02", birthDateRaw)
	if err != nil {
		return nil, ErrAuthInvalidInput
	}

	loginExists, err := s.authRepository.ExistsByLogin(login)
	if err != nil {
		return nil, err
	}

	if loginExists {
		return nil, ErrAuthLoginAlreadyUsed
	}

	emailExists, err := s.authRepository.ExistsByEmail(email)
	if err != nil {
		return nil, err
	}

	if emailExists {
		return nil, ErrAuthEmailAlreadyUsed
	}

	passwordHashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	response, err := s.authRepository.RegisterTouristUser(
		login,
		email,
		string(passwordHashBytes),
		firstName,
		lastName,
		middleName,
		sex,
		birthDate,
	)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrAuthTouristRoleNotFound
	}

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *AuthService) Login(request dto.LoginRequest) (*dto.LoginResponse, error) {
	login := strings.TrimSpace(request.Login)

	if login == "" || request.Password == "" {
		return nil, ErrAuthInvalidInput
	}

	user, err := s.authRepository.FindUserByLogin(login)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrAuthInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password)); err != nil {
		return nil, ErrAuthInvalidCredentials
	}

	token, err := s.createAccessToken(user.ID, user.Login)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
	}, nil
}

func (s *AuthService) Me(userID int64) (*dto.MeResponse, error) {
	if userID <= 0 {
		return nil, ErrAuthInvalidInput
	}

	response, err := s.authRepository.FindMeByUserID(userID)
	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, ErrAuthUserNotFound
	}

	return response, nil
}

func (s *AuthService) createAccessToken(userID int64, login string) (string, error) {
	ttlMinutes := s.jwtAccessTTLMinutes
	if ttlMinutes <= 0 {
		ttlMinutes = 60
	}

	now := time.Now()
	expiresAt := now.Add(time.Duration(ttlMinutes) * time.Minute)

	claims := jwt.MapClaims{
		"user_id": userID,
		"login":   login,
		"iat":     now.Unix(),
		"exp":     expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.jwtSecret))
}

func normalizeAuthStringPointer(value *string) *string {
	if value == nil {
		return nil
	}

	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		return nil
	}

	return &trimmed
}
