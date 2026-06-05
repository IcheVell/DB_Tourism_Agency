package models

import "time"

type TouristGroup struct {
	ID            int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ArrivalDate   time.Time `gorm:"column:arrival_date;not null" json:"arrival_date"`
	DepartureDate time.Time `gorm:"column:departure_date;not null" json:"departure_date"`
	Name          string    `gorm:"column:name;not null;unique" json:"name"`
}

func (TouristGroup) TableName() string {
	return "tourist_groups"
}
