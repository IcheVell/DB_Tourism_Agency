package dto

type CreateFlightRequest struct {
	Capacity     int    `json:"capacity"`
	FlightDate   string `json:"flight_date"`
	FlightTypeID int64  `json:"flight_type_id"`
}

type UpdateFlightRequest struct {
	Capacity     int    `json:"capacity"`
	FlightDate   string `json:"flight_date"`
	FlightTypeID int64  `json:"flight_type_id"`
}

type FlightResponse struct {
	ID           int64  `json:"id"`
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
