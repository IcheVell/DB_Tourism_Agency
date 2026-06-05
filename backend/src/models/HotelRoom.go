package models

type HotelRoom struct {
	ID         int64   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RoomNumber int     `gorm:"column:room_number;not null" json:"room_number"`
	Capacity   int     `gorm:"column:capacity;not null" json:"capacity"`
	Price      float64 `gorm:"column:price;not null" json:"price"`
	HotelID    int64   `gorm:"column:hotel_id;not null" json:"hotel_id"`

	Hotel *Hotel `gorm:"foreignKey:HotelID;references:ID" json:"hotel,omitempty"`
}

func (HotelRoom) TableName() string {
	return "hotel_rooms"
}
