package models_requests_posts

import "github.com/google/uuid"

type CreateUserRatingRequest struct {
	UserEvaluatorId uuid.UUID `json:"userEvaluatorId"`
	UserAvaluatedId uuid.UUID `json:"userEvaluatedId"`
	RatingTypeId    uuid.UUID `json:"ratingTypeId"`
	Reason          string    `json:"reason"`
}