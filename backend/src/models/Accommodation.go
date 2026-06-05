package models

import "time"

type Accommodation struct {
	ID            int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Status        string     `gorm:"column:status;not null" json:"status"`
	CheckInAt     time.Time  `gorm:"column:check_in_at;not null" json:"check_in_at"`
	CheckOutAt    *time.Time `gorm:"column:check_out_at;" json:"check_out_at"`
	GroupMemberID int64      `gorm:"column:group_member_id;not null" json:"group_member_id"`
	HotelRoomID   int64      `gorm:"column:hotel_room_id;not null" json:"hotel_room_id"`

	GroupMember *GroupMember `gorm:"foreignKey:GroupMemberID;references:ID" json:"group_member,omitempty"`
	HotelRoom   *HotelRoom   `gorm:"foreignKey:HotelRoomID;references:ID" json:"hotel_room,omitempty"`
}

func (Accommodation) TableName() string {
	return "accommodations"
}
