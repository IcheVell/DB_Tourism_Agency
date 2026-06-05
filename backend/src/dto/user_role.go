package dto

type UserBriefResponse struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
}

type UserRoleResponse struct {
	UserID int64 `json:"user_id"`
	RoleID int64 `json:"role_id"`

	User *UserBriefResponse `json:"user,omitempty"`
	Role *RoleBriefResponse `json:"role,omitempty"`
}

type CreateUserRoleRequest struct {
	UserID int64 `json:"user_id"`
	RoleID int64 `json:"role_id"`
}

type UpdateUserRoleRequest struct {
	UserID *int64 `json:"user_id,omitempty"`
	RoleID *int64 `json:"role_id,omitempty"`
}

type UserRoleListResponse struct {
	Items []UserRoleResponse `json:"items"`
	Total int64              `json:"total"`
	Page  int                `json:"page"`
	Limit int                `json:"limit"`
}
