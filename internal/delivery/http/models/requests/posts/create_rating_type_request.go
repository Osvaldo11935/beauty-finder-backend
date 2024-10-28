package models_requests_posts

type CreateRatingTypeRequest struct{
	Type string `json:"type"`
	Description string `json:"description"`
}