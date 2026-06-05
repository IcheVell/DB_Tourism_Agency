package models

import "time"

type FinancialOperation struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Amount      float64   `gorm:"column:amount;not null" json:"amount"`
	OperationAt time.Time `gorm:"column:operation_at;not null" json:"operation_at"`
	Description *string   `gorm:"column:description" json:"description,omitempty"`

	FinancialCategoryID int64 `gorm:"column:financial_category_id;not null" json:"financial_category_id"`

	FlightID            *int64 `gorm:"column:flight_id" json:"flight_id,omitempty"`
	VisaID              *int64 `gorm:"column:visa_id" json:"visa_id,omitempty"`
	ExcursionScheduleID *int64 `gorm:"column:excursion_schedule_id" json:"excursion_schedule_id,omitempty"`
	ExcursionBookingID  *int64 `gorm:"column:excursion_booking_id" json:"excursion_booking_id,omitempty"`
	CargoShipmentID     *int64 `gorm:"column:cargo_shipment_id" json:"cargo_shipment_id,omitempty"`
	CargoStatementID    *int64 `gorm:"column:cargo_statement_id" json:"cargo_statement_id,omitempty"`
	AccommodationID     *int64 `gorm:"column:accommodation_id" json:"accommodation_id,omitempty"`

	FinancialCategory *FinancialCategory `gorm:"foreignKey:FinancialCategoryID;references:ID" json:"financial_category,omitempty"`

	Flight            *Flight            `gorm:"foreignKey:FlightID;references:ID" json:"flight,omitempty"`
	Visa              *Visa              `gorm:"foreignKey:VisaID;references:ID" json:"visa,omitempty"`
	ExcursionSchedule *ExcursionSchedule `gorm:"foreignKey:ExcursionScheduleID;references:ID" json:"excursion_schedule,omitempty"`
	ExcursionBooking  *ExcursionBooking  `gorm:"foreignKey:ExcursionBookingID;references:ID" json:"excursion_booking,omitempty"`
	CargoShipment     *CargoShipment     `gorm:"foreignKey:CargoShipmentID;references:ID" json:"cargo_shipment,omitempty"`
	CargoStatement    *CargoStatement    `gorm:"foreignKey:CargoStatementID;references:ID" json:"cargo_statement,omitempty"`
	Accommodation     *Accommodation     `gorm:"foreignKey:AccommodationID;references:ID" json:"accommodation,omitempty"`
}

func (FinancialOperation) TableName() string {
	return "financial_operations"
}
