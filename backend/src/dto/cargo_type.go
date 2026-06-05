package dto

type CreateCargoTypeRequest struct {
	Name string `json:"name"`
}

type UpdateCargoTypeRequest struct {
	Name string `json:"name"`
}

type CargoTypeResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CargoTypeListResponse struct {
	Items    []CargoTypeResponse `json:"items"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"page_size"`
	Total    int64               `json:"total"`
}
