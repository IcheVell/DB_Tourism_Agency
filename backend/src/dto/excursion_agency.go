package dto

type CreateExcursionAgencyRequest struct {
	Name string `json:"name"`
}

type UpdateExcursionAgencyRequest struct {
	Name string `json:"name"`
}

type ExcursionAgencyResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ExcursionAgencyListResponse struct {
	Items    []ExcursionAgencyResponse `json:"items"`
	Page     int                       `json:"page"`
	PageSize int                       `json:"page_size"`
	Total    int64                     `json:"total"`
}
