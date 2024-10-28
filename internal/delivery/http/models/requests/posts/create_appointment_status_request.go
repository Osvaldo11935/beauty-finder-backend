package models_requests_posts

type CreateAppointmentStatusRequest struct{
	Type string `json:"type"`
	Description string `json:"description"`
}