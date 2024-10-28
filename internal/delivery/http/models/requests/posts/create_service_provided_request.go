package models_requests_posts

import "github.com/google/uuid"

type CreateServiceProvidedRequest struct {
	ServiceIds    []uuid.UUID `json:"serviceIds"`
}