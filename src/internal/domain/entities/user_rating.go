package entities

import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/primitives"

	"github.com/google/uuid"
)

type UserRating struct {
	primitives.BaseAuditableEntity
	UserEvaluatorId uuid.UUID `gorm:"column:UserEvaluatorId;" json:"userEvaluatorId"`
	UserAvaluatedId uuid.UUID `gorm:"column:UserEvaluatedId;" json:"userEvaluatedId"`
	RatingTypeId    uuid.UUID `gorm:"column:RatingTypeId;" json:"RatingTypeId"`
	Reason          string    `gorm:"column:Reason" json:"reason"`
	UserEvaluator   *User     `gorm:"-" json:"userEvaluator"`
	UserAvaluated   *User     `gorm:"-" json:"userAvaluated"`
}

func (s *UserRating) TableName() string {
	return "UserRating"
}

func NewUserRating(request models_requests_posts.CreateUserRatingRequest) UserRating {
	body := UserRating{
		BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
		UserEvaluatorId:     request.UserEvaluatorId,
		UserAvaluatedId:     request.UserAvaluatedId,
		RatingTypeId:        request.RatingTypeId,
	}

	return body
}

func (s *UserRating) Update(request models_requests_puts.UpdateUserRatingRequest) {
	if request.UserAvaluatedId != nil {
		s.UserAvaluatedId = *request.UserAvaluatedId
	}

	if request.UserEvaluatorId != nil {
		s.UserEvaluatorId = *request.UserEvaluatorId
	}

	if request.RatingTypeId != nil {
		s.RatingTypeId = *request.RatingTypeId
	}

	if request.Reason != nil {
		s.Reason = *request.Reason
	}

}
