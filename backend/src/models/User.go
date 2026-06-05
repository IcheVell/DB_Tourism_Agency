package models

import "time"

type User struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Login        string    `gorm:"column:login;not null;unique" json:"login"`
	PasswordHash string    `gorm:"column:password_hash;not null" json:"-"`
	Email        string    `gorm:"column:email;not null;unique" json:"email"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
