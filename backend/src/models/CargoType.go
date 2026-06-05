package models

type CargoType struct {
	ID   int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"column:name;not null;unique" json:"name"`
}

func (CargoType) TableName() string {
	return "cargo_types"
}
