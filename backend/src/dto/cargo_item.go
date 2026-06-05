package dto

type CreateCargoItemRequest struct {
	ItemNumber         string  `json:"item_number"`
	WeightKg           float64 `json:"weight_kg"`
	VolumetricWeightKg float64 `json:"volumetric_weight_kg"`
	PlacesCount        int     `json:"places_count"`
	Marking            *string `json:"marking,omitempty"`
	PackagedAt         *string `json:"packaged_at,omitempty"`
	CargoTypeID        int64   `json:"cargo_type_id"`
	CargoStatementID   int64   `json:"cargo_statement_id"`
}

type UpdateCargoItemRequest struct {
	ItemNumber         string  `json:"item_number"`
	WeightKg           float64 `json:"weight_kg"`
	VolumetricWeightKg float64 `json:"volumetric_weight_kg"`
	PlacesCount        int     `json:"places_count"`
	Marking            *string `json:"marking,omitempty"`
	PackagedAt         *string `json:"packaged_at,omitempty"`
	CargoTypeID        int64   `json:"cargo_type_id"`
	CargoStatementID   int64   `json:"cargo_statement_id"`
}

type CargoItemResponse struct {
	ID                 int64   `json:"id"`
	ItemNumber         string  `json:"item_number"`
	WeightKg           float64 `json:"weight_kg"`
	VolumetricWeightKg float64 `json:"volumetric_weight_kg"`
	PlacesCount        int     `json:"places_count"`
	Marking            *string `json:"marking,omitempty"`
	PackagedAt         *string `json:"packaged_at,omitempty"`
	CargoTypeID        int64   `json:"cargo_type_id"`
	CargoStatementID   int64   `json:"cargo_statement_id"`
}

type CargoItemListResponse struct {
	Items    []CargoItemResponse `json:"items"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"page_size"`
	Total    int64               `json:"total"`
}
