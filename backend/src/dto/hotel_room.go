package dto

type CreateHotelRoomRequest struct {
	RoomNumber int     `json:"room_number"`
	Capacity   int     `json:"capacity"`
	Price      float64 `json:"price"`
	HotelID    int64   `json:"hotel_id"`
}

type UpdateHotelRoomRequest struct {
	RoomNumber int     `json:"room_number"`
	Capacity   int     `json:"capacity"`
	Price      float64 `json:"price"`
	HotelID    int64   `json:"hotel_id"`
}

type HotelRoomResponse struct {
	ID         int64   `json:"id"`
	RoomNumber int     `json:"room_number"`
	Capacity   int     `json:"capacity"`
	Price      float64 `json:"price"`
	HotelID    int64   `json:"hotel_id"`
}

type HotelRoomListResponse struct {
	Items    []HotelRoomResponse `json:"items"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"page_size"`
	Total    int64               `json:"total"`
}
