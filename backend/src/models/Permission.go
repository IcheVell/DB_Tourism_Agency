package models

type Permission struct {
	ID          int64   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Code        string  `gorm:"column:code;not null;unique" json:"code"`
	Description *string `gorm:"column:description" json:"description,omitempty"`
}

func (Permission) TableName() string {
	return "permissions"
}
