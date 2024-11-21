package models_requests_posts

import "github.com/google/uuid"

type CreateUserRequest struct {
	Email       *string    `json:"email"`
	UserName    *string    `json:"userName"`
	Password    *string    `json:"password"`
	PhoneNumber string     `json:"phoneNumber"`
	CompanyId   *uuid.UUID `json:"companyId"`
}
