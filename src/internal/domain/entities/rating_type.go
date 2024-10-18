package entities

import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/primitives"
)

type RatingType struct {
	primitives.BaseAuditableEntity
	Type string `gorm:"column:Type;" json:"type"`
	Description string `gorm:"column:Description;" json:"description"`
	UserRating UserRating   `gorm:"foreignKey:RatingTypeId;references:Id" json:"userAvaluated"`
}

func(s *RatingType) TableName() string{
	return "RatingType"
}

func NewRatingType(request models_requests_posts.CreateRatingTypeRequest) RatingType{
	body := RatingType{
	   BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
	   Type: request.Type,
	   Description: request.Description,
   }

   return body
}

func(s *RatingType) Update(request models_requests_puts.UpdateRatingTypeRequest){
	if request.Type != nil {
		s.Type = *request.Type
	}

	if request.Description != nil {
       s.Description = *request.Description
	}
}