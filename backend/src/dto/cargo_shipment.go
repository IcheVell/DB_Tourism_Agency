package dto

type CreateCargoShipmentRequest struct {
	ShippedAt        *string `json:"shipped_at,omitempty"`
	Status           string  `json:"status"`
	CargoStatementID int64   `json:"cargo_statement_id"`
	FlightID         int64   `json:"flight_id"`
}

type UpdateCargoShipmentRequest struct {
	ShippedAt        *string `json:"shipped_at,omitempty"`
	Status           string  `json:"status"`
	CargoStatementID int64   `json:"cargo_statement_id"`
	FlightID         int64   `json:"flight_id"`
}

type CargoShipmentResponse struct {
	ID               int64   `json:"id"`
	ShippedAt        *string `json:"shipped_at,omitempty"`
	Status           string  `json:"status"`
	CargoStatementID int64   `json:"cargo_statement_id"`
	FlightID         int64   `json:"flight_id"`
}

type CargoShipmentListResponse struct {
	Items    []CargoShipmentResponse `json:"items"`
	Page     int                     `json:"page"`
	PageSize int                     `json:"page_size"`
	Total    int64                   `json:"total"`
}
