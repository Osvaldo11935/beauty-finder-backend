package models_responses

import (
	"src/internal/domain/entities"
	"github.com/google/uuid"
)

type UserRatingResponse struct {
	Id   uuid.UUID `json:"id"`
	UserEvaluator *UserResponse `json:"userEvaluator"`
	UserAvaluated *UserResponse `json:"userAvaluated"`
	RatingType   *RatingTypeResponse `json:"UserRating"`
	Reason          string    `json:"reason"`
}

func ToUserRatingResponse(data *entities.UserRating) *UserRatingResponse {
	if data == nil {
		return nil
	}
	return &UserRatingResponse{
		Id:   data.ID,
		Reason: data.Reason,
		RatingType: ToRatingTypeResponse(data.RatingType),
		UserEvaluator: ToUserResponse(data.UserEvaluator),
		UserAvaluated: ToUserResponse(data.UserAvaluated),
	}
}

func ToListUserRatingResponse(data []entities.UserRating) []UserRatingResponse {
	var resp []UserRatingResponse

	for _, item := range data {
		resp = append(resp, *ToUserRatingResponse(&item))
	}

	return resp
}
