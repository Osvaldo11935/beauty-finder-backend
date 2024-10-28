package models_requests_posts

type DispatchServiceNotification struct {
	RadiusKm float64  `json:"radiusKm"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
