package dto

type CreateFlightRequest struct {
	FlightNumber int    `json:"flight_number"`
	Capacity     int    `json:"capacity"`
	FlightDate   string `json:"flight_date"`
	FlightTypeID int64  `json:"flight_type_id"`
}

type UpdateFlightRequest struct {
	FlightNumber int    `json:"flight_number"`
	Capacity     int    `json:"capacity"`
	FlightDate   string `json:"flight_date"`
	FlightTypeID int64  `json:"flight_type_id"`
}

type FlightResponse struct {
	ID           int64  `json:"id"`
	FlightNumber int    `json:"flight_number"`
	Capacity     int    `json:"capacity"`
	FlightDate   string `json:"flight_date"`
	FlightTypeID int64  `json:"flight_type_id"`
}

type FlightListResponse struct {
	Items    []FlightResponse `json:"items"`
	Page     int              `json:"page"`
	PageSize int              `json:"page_size"`
	Total    int64            `json:"total"`
}
