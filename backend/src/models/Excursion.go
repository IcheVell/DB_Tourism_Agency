package models

type Excursion struct {
	ID          int64   `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	Name        string  `gorm:"column:name;not null" json:"name"`
	Description *string `gorm:"column:description" json:"description,omitempty"`
}

func (Excursion) TableName() string {
	return "excursions"
}
