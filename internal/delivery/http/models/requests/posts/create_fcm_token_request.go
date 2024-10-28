package models_requests_posts

type CreateFcmTokenRequest struct{
	FcmToken string `json:"fcmToken"`
	DeviceName string  `json:"deviceName"`
	DeviceId string `json:"deviceId"`
}