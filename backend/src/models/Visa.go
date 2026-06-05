package models

import "time"

type Visa struct {
	ID                 int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Number             *string    `gorm:"column:number" json:"number"`
	DestinationCountry string     `gorm:"column:destination_country;not null" json:"destination_country"`
	Status             string     `gorm:"column:status;not null" json:"status"`
	CreatedAt          time.Time  `gorm:"column:created_at;not null" json:"created_at"`
	SubmittedAt        *time.Time `gorm:"column:submitted_at" json:"submitted_at,omitempty"`
	DecisionAt         *time.Time `gorm:"column:decision_at" json:"decision_at,omitempty"`
	IssuedAt           *time.Time `gorm:"column:issued_at" json:"issued_at,omitempty"`
	ValidFrom          *time.Time `gorm:"column:valid_from" json:"valid_from,omitempty"`
	ValidUntil         *time.Time `gorm:"column:valid_until" json:"valid_until,omitempty"`
	TouristID          int64      `gorm:"column:tourist_id;not null" json:"tourist_id"`

	Tourist *Tourist `gorm:"foreignKey:TouristID;references:ID" json:"tourist,omitempty"`
}

func (Visa) TableName() string {
	return "visas"
}
