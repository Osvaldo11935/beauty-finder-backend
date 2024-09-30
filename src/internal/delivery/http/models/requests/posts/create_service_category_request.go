package models_requests_posts

type CreateServiceCategoryRequest struct{
	Name        string `json:"name"`
	Description string `json:"description"`
}