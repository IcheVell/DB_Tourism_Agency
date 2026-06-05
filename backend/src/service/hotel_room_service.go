package service

import (
	"errors"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"
)

var (
	ErrHotelRoomNotFound        = errors.New("hotel room not found")
	ErrHotelRoomInvalidNumber   = errors.New("hotel room number must be positive")
	ErrHotelRoomInvalidCapacity = errors.New("hotel room capacity must be positive")
	ErrHotelRoomInvalidPrice    = errors.New("hotel room price must be positive")
	ErrHotelRoomInvalidHotelID  = errors.New("hotel id must be positive")
)

type HotelRoomService struct {
	repo *repository.HotelRoomRepository
}

func NewHotelRoomService(repo *repository.HotelRoomRepository) *HotelRoomService {
	return &HotelRoomService{
		repo: repo,
	}
}

func (s *HotelRoomService) Create(req dto.CreateHotelRoomRequest) (*dto.HotelRoomResponse, error) {
	if req.RoomNumber <= 0 {
		return nil, ErrHotelRoomInvalidNumber
	}

	if req.Capacity <= 0 {
		return nil, ErrHotelRoomInvalidCapacity
	}

	if req.Price <= 0 {
		return nil, ErrHotelRoomInvalidPrice
	}

	if req.HotelID <= 0 {
		return nil, ErrHotelRoomInvalidHotelID
	}

	room := &models.HotelRoom{
		RoomNumber: req.RoomNumber,
		Capacity:   req.Capacity,
		Price:      req.Price,
		HotelID:    req.HotelID,
	}

	if err := s.repo.Create(room); err != nil {
		return nil, err
	}

	return toHotelRoomResponse(room), nil
}

func (s *HotelRoomService) GetByID(id int64) (*dto.HotelRoomResponse, error) {
	room, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if room == nil {
		return nil, ErrHotelRoomNotFound
	}

	return toHotelRoomResponse(room), nil
}

func (s *HotelRoomService) List(page int, pageSize int) (*dto.HotelRoomListResponse, error) {
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

	rooms, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.HotelRoomResponse, 0, len(rooms))
	for _, room := range rooms {
		items = append(items, *toHotelRoomResponse(&room))
	}

	return &dto.HotelRoomListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *HotelRoomService) Update(id int64, req dto.UpdateHotelRoomRequest) (*dto.HotelRoomResponse, error) {
	if req.RoomNumber <= 0 {
		return nil, ErrHotelRoomInvalidNumber
	}

	if req.Capacity <= 0 {
		return nil, ErrHotelRoomInvalidCapacity
	}

	if req.Price <= 0 {
		return nil, ErrHotelRoomInvalidPrice
	}

	if req.HotelID <= 0 {
		return nil, ErrHotelRoomInvalidHotelID
	}

	room, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if room == nil {
		return nil, ErrHotelRoomNotFound
	}

	room.RoomNumber = req.RoomNumber
	room.Capacity = req.Capacity
	room.Price = req.Price
	room.HotelID = req.HotelID

	if err := s.repo.Update(room); err != nil {
		return nil, err
	}

	return toHotelRoomResponse(room), nil
}

func (s *HotelRoomService) Delete(id int64) error {
	room, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if room == nil {
		return ErrHotelRoomNotFound
	}

	return s.repo.Delete(room)
}

func toHotelRoomResponse(room *models.HotelRoom) *dto.HotelRoomResponse {
	return &dto.HotelRoomResponse{
		ID:         room.ID,
		RoomNumber: room.RoomNumber,
		Capacity:   room.Capacity,
		Price:      room.Price,
		HotelID:    room.HotelID,
	}
}
