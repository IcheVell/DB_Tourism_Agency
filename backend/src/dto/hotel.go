package dto

type CreateHotelRequest struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type UpdateHotelRequest struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type HotelResponse struct {
	ID      int64  `json:"id"`
	Address string `json:"address"`
	Name    string `json:"name"`
}

type HotelListResponse struct {
	Items    []HotelResponse `json:"items"`
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
	Total    int64           `json:"total"`
}
