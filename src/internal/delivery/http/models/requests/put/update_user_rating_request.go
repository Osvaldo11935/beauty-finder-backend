package models_requests_puts

import "github.com/google/uuid"

type UpdateUserRatingRequest struct {
	UserEvaluatorId *uuid.UUID `json:"userEvaluatorId"`
	UserAvaluatedId *uuid.UUID `json:"userEvaluatedId"`
	RatingTypeId    *uuid.UUID `json:"RatingTypeId"`
	Reason          *string    `json:"reason"`
}