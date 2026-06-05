package dto

type RoleBriefResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type RolePermissionResponse struct {
	RoleID       int64               `json:"role_id"`
	PermissionID int64               `json:"permission_id"`
	Role         *RoleBriefResponse  `json:"role,omitempty"`
	Permission   *PermissionResponse `json:"permission,omitempty"`
}

type CreateRolePermissionRequest struct {
	RoleID       int64 `json:"role_id"`
	PermissionID int64 `json:"permission_id"`
}

type UpdateRolePermissionRequest struct {
	RoleID       *int64 `json:"role_id,omitempty"`
	PermissionID *int64 `json:"permission_id,omitempty"`
}

type RolePermissionListResponse struct {
	Items []RolePermissionResponse `json:"items"`
	Total int64                    `json:"total"`
	Page  int                      `json:"page"`
	Limit int                      `json:"limit"`
}
