package models_requests_posts

type CreateAttachmentTypeRequest struct{
	Type string `json:"type"`
	Description string `json:"description"`
}