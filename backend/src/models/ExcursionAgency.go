package models

type ExcursionAgency struct {
	ID   int64  `gorm:"column:id; primary_key;autoIncrement" json:"id"`
	Name string `gorm:"column:name;not null;unique" json:"name"`
}

func (ExcursionAgency) TableName() string {
	return "excursion_agencies"
}
