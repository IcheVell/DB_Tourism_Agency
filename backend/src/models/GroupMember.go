package models

type GroupMember struct {
	ID                int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TouristGroupID    int64  `gorm:"column:tourist_group_id;not null;" json:"tourist_group_id"`
	TouristCategoryID int64  `gorm:"column:tourist_category_id;not null" json:"tourist_category_id"`
	TouristID         int64  `gorm:"column:tourist_id;not null" json:"tourist_id"`
	DesiredHotelID    *int64 `gorm:"column:desired_hotel_id" json:"desired_hotel_id"`

	TouristGroup    *TouristGroup    `gorm:"foreignKey:TouristGroupID" json:"tourist_group,omitempty"`
	TouristCategory *TouristCategory `gorm:"foreignKey:TouristCategoryID" json:"tourist_category,omitempty"`
	Tourist         *Tourist         `gorm:"foreignKey:TouristID;references:ID" json:"tourist,omitempty"`
	Hotel           *Hotel           `gorm:"foreignKey:HotelID;references:ID" json:"hotel,omitempty"`
}

func (GroupMember) TableName() string {
	return "group_members"
}
