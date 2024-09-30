package models_responses

import "github.com/google/uuid"

type CreateResponse struct {
	Id uuid.UUID `json:"id"`
}

func NewCreateResponse(id uuid.UUID) CreateResponse{
	 return CreateResponse{Id: id}
}