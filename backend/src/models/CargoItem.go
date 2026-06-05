package models

import "time"

type CargoItem struct {
	ID                 int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	WeightKg           float64    `gorm:"column:weight_kg;not null" json:"weight_kg"`
	VolumetricWeightKg float64    `gorm:"column:volumetric_weight_kg;not null" json:"volumetric_weight_kg"`
	PlacesCount        int        `gorm:"column:places_count;not null" json:"places_count"`
	Marking            *string    `gorm:"column:marking" json:"marking,omitempty"`
	PackagedAt         *time.Time `gorm:"column:packaged_at" json:"packaged_at,omitempty"`
	ItemNumber         string     `gorm:"column:item_number;not null" json:"item_number"`

	CargoTypeID      int64 `gorm:"column:cargo_type_id;not null" json:"cargo_type_id"`
	CargoStatementID int64 `gorm:"column:cargo_statement_id;not null" json:"cargo_statement_id"`

	CargoType      *CargoType      `gorm:"foreignKey:CargoTypeID;references:ID" json:"cargo_type,omitempty"`
	CargoStatement *CargoStatement `gorm:"foreignKey:CargoStatementID;references:ID" json:"cargo_statement,omitempty"`
}

func (CargoItem) TableName() string {
	return "cargo_items"
}
