package models

import "time"

type ExcursionBooking struct {
	ID                  int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	BookedAt            time.Time `gorm:"column:booked_at;not null" json:"booked_at"`
	TouristRating       *int      `gorm:"column:tourist_rating" json:"tourist_rating,omitempty"`
	Status              string    `gorm:"column:status;not null" json:"status"`
	ExcursionScheduleID int64     `gorm:"column:excursion_schedule_id;not null" json:"excursion_schedule_id"`
	GroupMemberID       int64     `gorm:"column:group_member_id;not null" json:"group_member_id"`

	ExcursionSchedule *ExcursionSchedule `gorm:"foreignKey:ExcursionScheduleID;references:ID" json:"excursion_schedule,omitempty"`
	GroupMember       *GroupMember       `gorm:"foreignKey:GroupMemberID;references:ID" json:"group_member,omitempty"`
}

func (ExcursionBooking) TableName() string {
	return "excursion_bookings"
}
