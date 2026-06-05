package dto

type CreateExcursionScheduleRequest struct {
	Price             float64 `json:"price"`
	StartTime         string  `json:"start_time"`
	EndTime           string  `json:"end_time"`
	Capacity          int     `json:"capacity"`
	Status            string  `json:"status"`
	ExcursionAgencyID int64   `json:"excursion_agency_id"`
	ExcursionID       int64   `json:"excursion_id"`
}

type UpdateExcursionScheduleRequest struct {
	Price             float64 `json:"price"`
	StartTime         string  `json:"start_time"`
	EndTime           string  `json:"end_time"`
	Capacity          int     `json:"capacity"`
	Status            string  `json:"status"`
	ExcursionAgencyID int64   `json:"excursion_agency_id"`
	ExcursionID       int64   `json:"excursion_id"`
}

type ExcursionScheduleResponse struct {
	ID                int64   `json:"id"`
	Price             float64 `json:"price"`
	StartTime         string  `json:"start_time"`
	EndTime           string  `json:"end_time"`
	Capacity          int     `json:"capacity"`
	Status            string  `json:"status"`
	ExcursionAgencyID int64   `json:"excursion_agency_id"`
	ExcursionID       int64   `json:"excursion_id"`
}

type ExcursionScheduleListResponse struct {
	Items    []ExcursionScheduleResponse `json:"items"`
	Page     int                         `json:"page"`
	PageSize int                         `json:"page_size"`
	Total    int64                       `json:"total"`
}
