package models_requests_puts

import "github.com/google/uuid"

type UpdateUserRequest struct {
	Email       *string `json:"email"`
	Password    *string `json:"password"`
	PhoneNumber *string `json:"phoneNumber"`
	RoleId      *uuid.UUID `json:"roleId"`
}