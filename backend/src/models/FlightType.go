package models

type FlightType struct {
	ID   int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"column:name;not null;unique" json:"name"`
}

func (FlightType) TableName() string {
	return "flight_types"
}
