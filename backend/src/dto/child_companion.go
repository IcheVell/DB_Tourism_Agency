package dto

type CreateChildCompanionRequest struct {
	AdultGroupMemberID int64 `json:"adult_group_member_id"`
	ChildGroupMemberID int64 `json:"child_group_member_id"`
}

type UpdateChildCompanionRequest struct {
	AdultGroupMemberID int64 `json:"adult_group_member_id"`
	ChildGroupMemberID int64 `json:"child_group_member_id"`
}

type ChildCompanionResponse struct {
	AdultGroupMemberID int64 `json:"adult_group_member_id"`
	ChildGroupMemberID int64 `json:"child_group_member_id"`
}

type ChildCompanionListResponse struct {
	Items    []ChildCompanionResponse `json:"items"`
	Page     int                      `json:"page"`
	PageSize int                      `json:"page_size"`
	Total    int64                    `json:"total"`
}
