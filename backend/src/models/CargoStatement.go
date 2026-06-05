package models

type CargoStatement struct {
	ID            int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Status        string `gorm:"column:status" json:"status,omitempty"`
	GroupMemberID int64  `gorm:"column:group_member_id" json:"group_member_id,omitempty"`

	GroupMember *GroupMember `gorm:"foreignKey:GroupMemberID;references:ID" json:"group_member,omitempty"`
}

func (CargoStatement) TableName() string {
	return "cargo_statements"
}
