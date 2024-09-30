package models_requests_puts

type UpdateAttachmentTypeRequest struct{
	Type *string `json:"type"`
	Description *string `json:"description"`
}