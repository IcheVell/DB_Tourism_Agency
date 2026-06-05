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
	ErrTouristNotFound         = errors.New("tourist not found")
	ErrTouristInvalidFirstName = errors.New("tourist first name is required")
	ErrTouristInvalidLastName  = errors.New("tourist last name is required")
	ErrTouristInvalidSex       = errors.New("tourist sex must be male or female")
	ErrTouristInvalidBirthDate = errors.New("tourist birth date must be in YYYY-MM-DD format")
	ErrTouristInvalidUserID    = errors.New("user_id must be positive")
)

type TouristService struct {
	repo *repository.TouristRepository
}

func NewTouristService(repo *repository.TouristRepository) *TouristService {
	return &TouristService{
		repo: repo,
	}
}

func (s *TouristService) Create(req dto.CreateTouristRequest) (*dto.TouristResponse, error) {
	firstName := strings.TrimSpace(req.FirstName)
	lastName := strings.TrimSpace(req.LastName)
	sex := strings.TrimSpace(req.Sex)

	if firstName == "" {
		return nil, ErrTouristInvalidFirstName
	}

	if lastName == "" {
		return nil, ErrTouristInvalidLastName
	}

	if sex != "male" && sex != "female" {
		return nil, ErrTouristInvalidSex
	}

	birthDate, err := parseDateYYYYMMDD(req.BirthDate)
	if err != nil {
		return nil, ErrTouristInvalidBirthDate
	}

	if req.UserID != nil && *req.UserID <= 0 {
		return nil, ErrTouristInvalidUserID
	}

	tourist := &models.Tourist{
		FirstName:  firstName,
		LastName:   lastName,
		MiddleName: normalizeOptionalString(req.MiddleName),
		Sex:        sex,
		BirthDate:  birthDate,
		UserID:     req.UserID,
	}

	if err := s.repo.Create(tourist); err != nil {
		return nil, err
	}

	return toTouristResponse(tourist), nil
}

func (s *TouristService) GetByID(id int64) (*dto.TouristResponse, error) {
	tourist, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if tourist == nil {
		return nil, ErrTouristNotFound
	}

	return toTouristResponse(tourist), nil
}

func (s *TouristService) List(page int, pageSize int) (*dto.TouristListResponse, error) {
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

	tourists, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.TouristResponse, 0, len(tourists))
	for _, tourist := range tourists {
		items = append(items, *toTouristResponse(&tourist))
	}

	return &dto.TouristListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *TouristService) Update(id int64, req dto.UpdateTouristRequest) (*dto.TouristResponse, error) {
	firstName := strings.TrimSpace(req.FirstName)
	lastName := strings.TrimSpace(req.LastName)
	sex := strings.TrimSpace(req.Sex)

	if firstName == "" {
		return nil, ErrTouristInvalidFirstName
	}

	if lastName == "" {
		return nil, ErrTouristInvalidLastName
	}

	if sex != "male" && sex != "female" {
		return nil, ErrTouristInvalidSex
	}

	birthDate, err := parseDateYYYYMMDD(req.BirthDate)
	if err != nil {
		return nil, ErrTouristInvalidBirthDate
	}

	if req.UserID != nil && *req.UserID <= 0 {
		return nil, ErrTouristInvalidUserID
	}

	tourist, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if tourist == nil {
		return nil, ErrTouristNotFound
	}

	tourist.FirstName = firstName
	tourist.LastName = lastName
	tourist.MiddleName = normalizeOptionalString(req.MiddleName)
	tourist.Sex = sex
	tourist.BirthDate = birthDate
	tourist.UserID = req.UserID

	if err := s.repo.Update(tourist); err != nil {
		return nil, err
	}

	return toTouristResponse(tourist), nil
}

func (s *TouristService) Delete(id int64) error {
	tourist, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if tourist == nil {
		return ErrTouristNotFound
	}

	return s.repo.Delete(tourist)
}

func toTouristResponse(tourist *models.Tourist) *dto.TouristResponse {
	return &dto.TouristResponse{
		ID:         tourist.ID,
		FirstName:  tourist.FirstName,
		LastName:   tourist.LastName,
		MiddleName: tourist.MiddleName,
		Sex:        tourist.Sex,
		BirthDate:  tourist.BirthDate.Format("2006-01-02"),
		UserID:     tourist.UserID,
	}
}

func parseDateYYYYMMDD(value string) (time.Time, error) {
	return time.Parse("2006-01-02", strings.TrimSpace(value))
}
