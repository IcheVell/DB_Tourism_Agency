package models

import "time"

type Tourist struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	FirstName  string    `gorm:"column:first_name;not null" json:"first_name"`
	LastName   string    `gorm:"column:last_name;not null" json:"last_name"`
	MiddleName *string   `gorm:"column:middle_name" json:"middle_name,omitempty"`
	Sex        string    `gorm:"column:sex;not null" json:"sex"`
	BirthDate  time.Time `gorm:"column:birth_date;type:date;not null" json:"birth_date"`
	UserID     *int64    `gorm:"column:user_id" json:"user_id,omitempty"`

	User *User `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
}

func (Tourist) TableName() string {
	return "tourists"
}
