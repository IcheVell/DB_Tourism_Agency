package models

type RolePermission struct {
	RoleID       int64 `gorm:"column:role_id;primaryKey;not null" json:"role_id,omitempty"`
	PermissionID int64 `gorm:"column:permission_id;primaryKey;not null" json:"permission_id,omitempty"`

	Role       *Role       `gorm:"foreignKey:RoleID;references:ID" json:"role,omitempty"`
	Permission *Permission `gorm:"foreignKey:PermissionID;references:ID" json:"permission,omitempty"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
