package dto

import "time"

type CreateFinancialOperationRequest struct {
	Amount      float64    `json:"amount"`
	OperationAt *time.Time `json:"operation_at,omitempty"`
	Description *string    `json:"description,omitempty"`

	FinancialCategoryID int64 `json:"financial_category_id"`

	FlightID            *int64 `json:"flight_id,omitempty"`
	VisaID              *int64 `json:"visa_id,omitempty"`
	ExcursionScheduleID *int64 `json:"excursion_schedule_id,omitempty"`
	ExcursionBookingID  *int64 `json:"excursion_booking_id,omitempty"`
	CargoShipmentID     *int64 `json:"cargo_shipment_id,omitempty"`
	CargoStatementID    *int64 `json:"cargo_statement_id,omitempty"`
	AccommodationID     *int64 `json:"accommodation_id,omitempty"`
}

type UpdateFinancialOperationRequest struct {
	Amount      float64    `json:"amount"`
	OperationAt *time.Time `json:"operation_at,omitempty"`
	Description *string    `json:"description,omitempty"`

	FinancialCategoryID int64 `json:"financial_category_id"`

	FlightID            *int64 `json:"flight_id,omitempty"`
	VisaID              *int64 `json:"visa_id,omitempty"`
	ExcursionScheduleID *int64 `json:"excursion_schedule_id,omitempty"`
	ExcursionBookingID  *int64 `json:"excursion_booking_id,omitempty"`
	CargoShipmentID     *int64 `json:"cargo_shipment_id,omitempty"`
	CargoStatementID    *int64 `json:"cargo_statement_id,omitempty"`
	AccommodationID     *int64 `json:"accommodation_id,omitempty"`
}

type FinancialOperationResponse struct {
	ID          int64     `json:"id"`
	Amount      float64   `json:"amount"`
	OperationAt time.Time `json:"operation_at"`
	Description *string   `json:"description,omitempty"`

	FinancialCategoryID int64 `json:"financial_category_id"`

	FlightID            *int64 `json:"flight_id,omitempty"`
	VisaID              *int64 `json:"visa_id,omitempty"`
	ExcursionScheduleID *int64 `json:"excursion_schedule_id,omitempty"`
	ExcursionBookingID  *int64 `json:"excursion_booking_id,omitempty"`
	CargoShipmentID     *int64 `json:"cargo_shipment_id,omitempty"`
	CargoStatementID    *int64 `json:"cargo_statement_id,omitempty"`
	AccommodationID     *int64 `json:"accommodation_id,omitempty"`
}

type FinancialOperationListResponse struct {
	Items    []FinancialOperationResponse `json:"items"`
	Page     int                          `json:"page"`
	PageSize int                          `json:"page_size"`
	Total    int64                        `json:"total"`
}
