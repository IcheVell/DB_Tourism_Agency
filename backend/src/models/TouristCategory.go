package models

type TouristCategory struct {
	ID   int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"column:name;not null;unique" json:"name"`
}

func (TouristCategory) TableName() string {
	return "tourist_categories"
}
