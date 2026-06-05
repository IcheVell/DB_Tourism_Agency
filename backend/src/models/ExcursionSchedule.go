package models

import "time"

type ExcursionSchedule struct {
	ID                int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Price             float64   `gorm:"column:price;not null" json:"price"`
	StartTime         time.Time `gorm:"column:start_time;not null" json:"start_time"`
	EndTime           time.Time `gorm:"column:end_time;not null" json:"end_time"`
	Capacity          int       `gorm:"column:capacity;not null" json:"capacity"`
	Status            string    `gorm:"column:status;not null" json:"status"`
	ExcursionAgencyID int64     `gorm:"column:excursion_agency_id;not null" json:"excursion_agency_id"`
	ExcursionID       int64     `gorm:"column:excursion_id;not null" json:"excursion_id"`

	ExcursionAgency *ExcursionAgency `gorm:"foreignKey:ExcursionAgencyID;references:ID" json:"excursion_agency,omitempty"`
	Excursion       *Excursion       `gorm:"foreignKey:ExcursionID;references:ID" json:"excursion,omitempty"`
}

func (ExcursionSchedule) TableName() string {
	return "excursion_schedule"
}
