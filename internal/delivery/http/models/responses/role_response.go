package models_responses

import (
	"src/internal/domain/entities"

	"github.com/google/uuid"
)

type RoleResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func ToRoleResponse(data *entities.Role) *RoleResponse {
	if data == nil {
		return nil
	}
	return &RoleResponse{
		Id:   data.ID,
		Name: data.Name,
	}
}

func ToListRoleResponse(data []entities.Role) []RoleResponse {
	var resp []RoleResponse

	for _, item := range data {
		resp = append(resp, *ToRoleResponse(&item))
	}

	return resp
}
