package dto

type CreateTouristRequest struct {
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	MiddleName     *string `json:"middle_name,omitempty"`
	Sex            string  `json:"sex"`
	BirthDate      string  `json:"birth_date"`
	UserID         *int64  `json:"user_id,omitempty"`
	DesiredHotelID *int64  `json:"desired_hotel_id,omitempty"`
}

type UpdateTouristRequest struct {
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	MiddleName     *string `json:"middle_name,omitempty"`
	Sex            string  `json:"sex"`
	BirthDate      string  `json:"birth_date"`
	UserID         *int64  `json:"user_id,omitempty"`
	DesiredHotelID *int64  `json:"desired_hotel_id,omitempty"`
}

type TouristResponse struct {
	ID             int64   `json:"id"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	MiddleName     *string `json:"middle_name,omitempty"`
	Sex            string  `json:"sex"`
	BirthDate      string  `json:"birth_date"`
	UserID         *int64  `json:"user_id,omitempty"`
	DesiredHotelID *int64  `json:"desired_hotel_id,omitempty"`
}

type TouristListResponse struct {
	Items    []TouristResponse `json:"items"`
	Page     int               `json:"page"`
	PageSize int               `json:"page_size"`
	Total    int64             `json:"total"`
}
