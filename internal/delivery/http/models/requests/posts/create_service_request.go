package models_requests_posts

import "github.com/google/uuid"

type CreateServiceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  uuid.UUID `json:"categoryId"`
}