package service

import (
	"errors"
	"strings"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"
)

var (
	ErrHotelNotFound       = errors.New("hotel not found")
	ErrHotelInvalidName    = errors.New("hotel name is required")
	ErrHotelInvalidAddress = errors.New("hotel address is required")
)

type HotelService struct {
	repo *repository.HotelRepository
}

func NewHotelService(repo *repository.HotelRepository) *HotelService {
	return &HotelService{
		repo: repo,
	}
}

func (s *HotelService) Create(req dto.CreateHotelRequest) (*dto.HotelResponse, error) {
	name := strings.TrimSpace(req.Name)
	address := strings.TrimSpace(req.Address)

	if name == "" {
		return nil, ErrHotelInvalidName
	}

	if address == "" {
		return nil, ErrHotelInvalidAddress
	}

	hotel := &models.Hotel{
		Name:    name,
		Address: address,
	}

	if err := s.repo.Create(hotel); err != nil {
		return nil, err
	}

	return toHotelResponse(hotel), nil
}

func (s *HotelService) GetByID(id int64) (*dto.HotelResponse, error) {
	hotel, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if hotel == nil {
		return nil, ErrHotelNotFound
	}

	return toHotelResponse(hotel), nil
}

func (s *HotelService) List(page int, pageSize int) (*dto.HotelListResponse, error) {
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

	hotels, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.HotelResponse, 0, len(hotels))
	for _, hotel := range hotels {
		items = append(items, *toHotelResponse(&hotel))
	}

	return &dto.HotelListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *HotelService) Update(id int64, req dto.UpdateHotelRequest) (*dto.HotelResponse, error) {
	name := strings.TrimSpace(req.Name)
	address := strings.TrimSpace(req.Address)

	if name == "" {
		return nil, ErrHotelInvalidName
	}

	if address == "" {
		return nil, ErrHotelInvalidAddress
	}

	hotel, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if hotel == nil {
		return nil, ErrHotelNotFound
	}

	hotel.Name = name
	hotel.Address = address

	if err := s.repo.Update(hotel); err != nil {
		return nil, err
	}

	return toHotelResponse(hotel), nil
}

func (s *HotelService) Delete(id int64) error {
	hotel, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if hotel == nil {
		return ErrHotelNotFound
	}

	return s.repo.Delete(hotel)
}

func toHotelResponse(hotel *models.Hotel) *dto.HotelResponse {
	return &dto.HotelResponse{
		ID:      hotel.ID,
		Address: hotel.Address,
		Name:    hotel.Name,
	}
}
