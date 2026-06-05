package dto

type CreateTouristGroupRequest struct {
	Name          string `json:"name"`
	ArrivalDate   string `json:"arrival_date"`
	DepartureDate string `json:"departure_date"`
}

type UpdateTouristGroupRequest struct {
	Name          string `json:"name"`
	ArrivalDate   string `json:"arrival_date"`
	DepartureDate string `json:"departure_date"`
}

type TouristGroupResponse struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	ArrivalDate   string `json:"arrival_date"`
	DepartureDate string `json:"departure_date"`
}

type TouristGroupListResponse struct {
	Items    []TouristGroupResponse `json:"items"`
	Page     int                    `json:"page"`
	PageSize int                    `json:"page_size"`
	Total    int64                  `json:"total"`
}
