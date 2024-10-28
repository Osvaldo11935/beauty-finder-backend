package models_requests_puts

type UpdateServiceCategoryRequest struct{
	Name        *string `json:"name"`
	Description *string `json:"description"`
}