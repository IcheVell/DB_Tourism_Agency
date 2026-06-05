package dto

import "time"

type UserResponse struct {
	ID        uint64        `json:"id"`
	Login     string        `json:"login"`
	Email     string        `json:"email"`
	Role      *RoleResponse `json:"role,omitempty"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type CreateUserRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   uint64 `json:"role_id"`
}

type UpdateUserRequest struct {
	Login    *string `json:"login,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	RoleID   *uint64 `json:"role_id,omitempty"`
}

type UserListResponse struct {
	Items []UserResponse `json:"items"`
	Total int64          `json:"total"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
