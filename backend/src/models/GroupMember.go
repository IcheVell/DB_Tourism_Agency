package models

type GroupMember struct {
	ID                int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TouristGroupID    int64  `gorm:"column:tourist_group_id;not null" json:"tourist_group_id"`
	TouristCategoryID int64  `gorm:"column:tourist_category_id;not null" json:"tourist_category_id"`
	TouristID         int64  `gorm:"column:tourist_id;not null" json:"tourist_id"`
	DesiredHotelID    *int64 `gorm:"column:desired_hotel_id" json:"desired_hotel_id,omitempty"`

	TouristGroup    *TouristGroup    `gorm:"foreignKey:TouristGroupID;references:ID" json:"tourist_group,omitempty"`
	TouristCategory *TouristCategory `gorm:"foreignKey:TouristCategoryID;references:ID" json:"tourist_category,omitempty"`
	Tourist         *Tourist         `gorm:"foreignKey:TouristID;references:ID" json:"tourist,omitempty"`
	DesiredHotel    *Hotel           `gorm:"foreignKey:DesiredHotelID;references:ID" json:"desired_hotel,omitempty"`
}

func (GroupMember) TableName() string {
	return "group_members"
}
