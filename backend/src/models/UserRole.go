package models

type UserRole struct {
	UserID int64 `gorm:"column:user_id;primaryKey" json:"user_id"`
	RoleID int64 `gorm:"column:role_id;primaryKey" json:"role_id"`

	User *User `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Role *Role `gorm:"foreignKey:RoleID;references:ID" json:"role,omitempty"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
