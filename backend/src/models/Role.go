package models

type Role struct {
	ID          int64   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name        string  `gorm:"column:name;not null;unique" json:"name"`
	Description *string `gorm:"column:description" json:"description,omitempty"`
}

func (Role) TableName() string {
	return "roles"
}
