package dto

type RoleResponse struct {
	ID          int64                `json:"id"`
	Name        string               `json:"name"`
	Description *string              `json:"description,omitempty"`
	Permissions []PermissionResponse `json:"permissions"`
}

type CreateRoleRequest struct {
	Name          string  `json:"name"`
	Description   *string `json:"description,omitempty"`
	PermissionIDs []int64 `json:"permission_ids"`
}

type UpdateRoleRequest struct {
	Name          *string  `json:"name,omitempty"`
	Description   *string  `json:"description,omitempty"`
	PermissionIDs *[]int64 `json:"permission_ids,omitempty"`
}

type RoleListResponse struct {
	Items []RoleResponse `json:"items"`
	Total int64          `json:"total"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
}
