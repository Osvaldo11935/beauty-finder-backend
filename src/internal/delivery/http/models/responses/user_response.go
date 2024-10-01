package models_responses

import (
	"src/internal/domain/entities"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type UserResponse struct {
	Id               uuid.UUID          `json:"id"`
	Email            string             `json:"Email"`
	UserName         string             `json:"UserName"`
	Password         string             `json:"Password"`
	PhoneNumber      string             `json:"PhoneNumber"`
	Person           *PersonResponse    `json:"person"`
	Conn             *websocket.Conn    `json:"conn"`
	Role             *RoleResponse `json:"role"`
}

func ToUserResponse(data *entities.User) *UserResponse {

	if data == nil {
		return nil
	}

	return &UserResponse{
		Id:          data.ID,
		Email:       data.Email,
		UserName:    data.UserName,
		PhoneNumber: data.PhoneNumber,
		Person:      ToPersonResponse(data.Person),
		Conn:        data.Conn,
		Role:        ToRoleResponse(data.Role),
	}
}

func ToListUserResponse(data []entities.User) []UserResponse {
	var resp []UserResponse

	for _, item := range data {
		resp = append(resp, *ToUserResponse(&item))
	}

	return resp
}