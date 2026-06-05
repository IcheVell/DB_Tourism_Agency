package dto

type CreateTouristCategoryRequest struct {
	Name string `json:"name"`
}

type UpdateTouristCategoryRequest struct {
	Name string `json:"name"`
}

type TouristCategoryResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type TouristCategoryListResponse struct {
	Items    []TouristCategoryResponse `json:"items"`
	Page     int                       `json:"page"`
	PageSize int                       `json:"page_size"`
	Total    int64                     `json:"total"`
}
