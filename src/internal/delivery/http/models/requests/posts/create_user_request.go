package models_requests_posts

type CreateUserRequest struct{
	Email    string `json:"email"`
	Password string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}