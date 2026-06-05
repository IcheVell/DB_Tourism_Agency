package dto

type CreateGroupMemberRequest struct {
	TouristGroupID    int64  `json:"tourist_group_id"`
	TouristCategoryID int64  `json:"tourist_category_id"`
	TouristID         int64  `json:"tourist_id"`
	DesiredHotelID    *int64 `json:"desired_hotel_id,omitempty"`
}

type UpdateGroupMemberRequest struct {
	TouristGroupID    int64  `json:"tourist_group_id"`
	TouristCategoryID int64  `json:"tourist_category_id"`
	TouristID         int64  `json:"tourist_id"`
	DesiredHotelID    *int64 `json:"desired_hotel_id,omitempty"`
}

type GroupMemberResponse struct {
	ID                int64  `json:"id"`
	TouristGroupID    int64  `json:"tourist_group_id"`
	TouristCategoryID int64  `json:"tourist_category_id"`
	TouristID         int64  `json:"tourist_id"`
	DesiredHotelID    *int64 `json:"desired_hotel_id,omitempty"`
}

type GroupMemberListResponse struct {
	Items    []GroupMemberResponse `json:"items"`
	Page     int                   `json:"page"`
	PageSize int                   `json:"page_size"`
	Total    int64                 `json:"total"`
}
