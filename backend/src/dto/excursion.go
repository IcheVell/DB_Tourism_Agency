package dto

type CreateExcursionRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type UpdateExcursionRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type ExcursionResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type ExcursionListResponse struct {
	Items    []ExcursionResponse `json:"items"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"page_size"`
	Total    int64               `json:"total"`
}
