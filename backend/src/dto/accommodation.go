package dto

type CreateAccommodationRequest struct {
	Status        string  `json:"status"`
	CheckInAt     string  `json:"check_in_at"`
	CheckOutAt    *string `json:"check_out_at,omitempty"`
	GroupMemberID int64   `json:"group_member_id"`
	HotelRoomID   int64   `json:"hotel_room_id"`
}

type UpdateAccommodationRequest struct {
	Status        string  `json:"status"`
	CheckInAt     string  `json:"check_in_at"`
	CheckOutAt    *string `json:"check_out_at,omitempty"`
	GroupMemberID int64   `json:"group_member_id"`
	HotelRoomID   int64   `json:"hotel_room_id"`
}

type AccommodationResponse struct {
	ID            int64   `json:"id"`
	Status        string  `json:"status"`
	CheckInAt     string  `json:"check_in_at"`
	CheckOutAt    *string `json:"check_out_at,omitempty"`
	GroupMemberID int64   `json:"group_member_id"`
	HotelRoomID   int64   `json:"hotel_room_id"`
}

type AccommodationListResponse struct {
	Items    []AccommodationResponse `json:"items"`
	Page     int                     `json:"page"`
	PageSize int                     `json:"page_size"`
	Total    int64                   `json:"total"`
}
