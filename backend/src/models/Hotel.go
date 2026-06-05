package models

type Hotel struct {
	ID      int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Address string `gorm:"column:address;not null" json:"address"`
	Name    string `gorm:"column:name;not null" json:"name"`
}

func (Hotel) TableName() string {
	return "hotels"
}
