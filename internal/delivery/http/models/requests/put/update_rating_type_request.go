package models_requests_puts

type UpdateRatingTypeRequest struct{
	Type *string `json:"type"`
	Description *string `json:"description"`
}