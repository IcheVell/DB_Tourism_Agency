package dto

type CreateFlightTypeRequest struct {
	Name string `json:"name"`
}

type UpdateFlightTypeRequest struct {
	Name string `json:"name"`
}

type FlightTypeResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type FlightTypeListResponse struct {
	Items    []FlightTypeResponse `json:"items"`
	Page     int                  `json:"page"`
	PageSize int                  `json:"page_size"`
	Total    int64                `json:"total"`
}
