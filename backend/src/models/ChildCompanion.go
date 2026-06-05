package models

type ChildCompanion struct {
	ChildGroupMemberID int64 `gorm:"column:child_group_member_id;primaryKey" json:"child_group_member_id"`
	AdultGroupMemberID int64 `gorm:"column:adult_group_member_id;primaryKey" json:"adult_group_member_id"`

	ChildGroupMember *GroupMember `gorm:"foreignKey:ChildGroupMemberID;references:ID" json:"child_group_member,omitempty"`
	AdultGroupMember *GroupMember `gorm:"foreignKey:AdultGroupMemberID;references:ID" json:"adult_group_member,omitempty"`
}

func (ChildCompanion) TableName() string {
	return "child_companions"
}
