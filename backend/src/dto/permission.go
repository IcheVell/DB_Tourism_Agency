package dto

type PermissionResponse struct {
	ID          int64   `json:"id"`
	Code        string  `json:"code"`
	Description *string `json:"description,omitempty"`
}

type CreatePermissionRequest struct {
	Code        string  `json:"code"`
	Description *string `json:"description,omitempty"`
}

type UpdatePermissionRequest struct {
	Code        *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

type PermissionListResponse struct {
	Items []PermissionResponse `json:"items"`
	Total int64                `json:"total"`
	Page  int                  `json:"page"`
	Limit int                  `json:"limit"`
}
