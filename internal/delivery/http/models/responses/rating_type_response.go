package models_responses

import (
	"src/internal/domain/entities"
	"github.com/google/uuid"
)

type RatingTypeResponse struct {
	Id   uuid.UUID `json:"id"`
	Type string    `json:"type"`
	Description string `json:"description"`
}

func ToRatingTypeResponse(data *entities.RatingType) *RatingTypeResponse {
	if data == nil {
		return nil
	}
	return &RatingTypeResponse{
		Id:   data.ID,
		Type: data.Type,
		Description: data.Description,
	}
}

func ToListRatingTypeResponse(data []entities.RatingType) []RatingTypeResponse {
	var resp []RatingTypeResponse

	for _, item := range data {
		resp = append(resp, *ToRatingTypeResponse(&item))
	}

	return resp
}
