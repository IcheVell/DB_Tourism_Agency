package dto

type CreateCargoStatementRequest struct {
	Status        string `json:"status"`
	GroupMemberID int64  `json:"group_member_id"`
}

type UpdateCargoStatementRequest struct {
	Status        string `json:"status"`
	GroupMemberID int64  `json:"group_member_id"`
}

type CargoStatementResponse struct {
	ID            int64  `json:"id"`
	Status        string `json:"status"`
	GroupMemberID int64  `json:"group_member_id"`
}

type CargoStatementListResponse struct {
	Items    []CargoStatementResponse `json:"items"`
	Page     int                      `json:"page"`
	PageSize int                      `json:"page_size"`
	Total    int64                    `json:"total"`
}
