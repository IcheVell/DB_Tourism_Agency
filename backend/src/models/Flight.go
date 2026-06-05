package models

import "time"

type Flight struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Capacity     int       `gorm:"column:capacity;not null" json:"capacity"`
	FlightDate   time.Time `gorm:"column:flight_date;not null" json:"flight_date"`
	FlightNumber int       `gorm:"column:flight_number;not null;unique" json:"flight_number"`
	FlightTypeID int64     `gorm:"column:flight_type_id" json:"flight_type_id,omitempty"`

	FlightType *FlightType `gorm:"foreignKey:FlightTypeID;references:ID" json:"flight_type,omitempty"`
}

func (Flight) TableName() string {
	return "flights"
}
