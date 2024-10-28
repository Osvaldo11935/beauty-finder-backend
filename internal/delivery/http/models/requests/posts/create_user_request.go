package models_requests_posts

type CreateUserRequest struct {
	Email       *string `json:"email"`
	UserName    *string `json:"userName"`
	Password    *string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}
