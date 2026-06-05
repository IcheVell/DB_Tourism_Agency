package dto

type CreateExcursionBookingRequest struct {
	BookedAt            string `json:"booked_at"`
	TouristRating       *int   `json:"tourist_rating,omitempty"`
	Status              string `json:"status"`
	ExcursionScheduleID int64  `json:"excursion_schedule_id"`
	GroupMemberID       int64  `json:"group_member_id"`
}

type UpdateExcursionBookingRequest struct {
	BookedAt            string `json:"booked_at"`
	TouristRating       *int   `json:"tourist_rating,omitempty"`
	Status              string `json:"status"`
	ExcursionScheduleID int64  `json:"excursion_schedule_id"`
	GroupMemberID       int64  `json:"group_member_id"`
}

type ExcursionBookingResponse struct {
	ID                  int64  `json:"id"`
	BookedAt            string `json:"booked_at"`
	TouristRating       *int   `json:"tourist_rating,omitempty"`
	Status              string `json:"status"`
	ExcursionScheduleID int64  `json:"excursion_schedule_id"`
	GroupMemberID       int64  `json:"group_member_id"`
}

type ExcursionBookingListResponse struct {
	Items    []ExcursionBookingResponse `json:"items"`
	Page     int                        `json:"page"`
	PageSize int                        `json:"page_size"`
	Total    int64                      `json:"total"`
}
