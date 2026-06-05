package models

import "time"

type CargoShipment struct {
	ID               int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ShippedAt        *time.Time `gorm:"column:shipped_at" json:"shipped_at,omitempty"`
	Status           string     `gorm:"column:status" json:"status,omitempty"`
	FlightID         int64      `gorm:"column:flight_id" json:"flight_id,omitempty"`
	CargoStatementID int64      `gorm:"column:cargo_statement_id" json:"cargo_statement_id,omitempty"`

	Flight         *Flight         `gorm:"foreignKey:FlightID;references:ID" json:"flight,omitempty"`
	CargoStatement *CargoStatement `gorm:"foreignKey:CargoStatementID;references:ID" json:"cargo_statement,omitempty"`
}

func (CargoShipment) TableName() string {
	return "cargo_shipments"
}
