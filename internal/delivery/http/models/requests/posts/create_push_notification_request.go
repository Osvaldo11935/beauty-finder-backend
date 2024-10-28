package models_requests_posts

type CreatePushNotificationRequest struct {
	To           string `json:"to"`
	Notification struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	} `json:"notification"`
	Data struct {
		ClickAction string `json:"click_action"`
	} `json:"data"`
}