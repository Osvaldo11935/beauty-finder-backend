package entities

import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/primitives"

	"github.com/google/uuid"
)

type Address struct {
	primitives.BaseAuditableEntity
	District     string  `gorm:"column:District;"`
	Commune      string  `gorm:"column:Commune;"`
	Province     string  `gorm:"column:Province;"`
	Country      string  `gorm:"column:Country;"`
	City         string  `gorm:"column:City;"`
	Street       string  `gorm:"column:Street;"`
	Neighborhood string  `gorm:"column:Neighborhood;"`
	Latitude     float64 `gorm:"column:Latitude;"`
	Longitude    float64 `gorm:"column:Longitude;"`
	UserId       uuid.UUID  `gorm:"column:UserId;"`
}

func (s *Address) TableName() string {
	return "Address"
}

func NewAddress(userId uuid.UUID, request models_requests_posts.CreateAddressRequest) Address {
	body := Address{
		BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
		District:            request.District,
		Commune:             request.Commune,
		Province:            request.Province,
		Country:             request.Country,
		City:                request.City,
		Street:              request.Street,
		Neighborhood:        request.Neighborhood,
		Latitude:            request.Latitude,
		Longitude:           request.Longitude,
		UserId:              userId,
	}

	return body
}

func (s *Address) Update(request models_requests_puts.UpdateAddressRequest) {
	if request.City != nil {
		s.City = *request.City
	}
	if request.District != nil {
		s.District = *request.District
	}
	if request.Commune != nil {
		s.Commune = *request.Commune
	}
	if request.Province != nil {
		s.Province = *request.Province
	}
	if request.Country != nil {
		s.Country = *request.Country
	}
	if request.Street != nil {
		s.Street = *request.Street
	}
	if request.Neighborhood != nil {
		s.Neighborhood = *request.Neighborhood
	}
	if request.Latitude != nil {
		s.Latitude = *request.Latitude
	}
	if request.Longitude != nil {
		s.Longitude = *request.Longitude
	}
}
