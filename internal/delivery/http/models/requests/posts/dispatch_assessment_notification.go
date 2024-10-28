package models_requests_posts

import "github.com/google/uuid"

type DispatchAssessmentNotification struct {
	
	UserAvaluatedId *uuid.UUID `json:"userAvaluatedId"`
}
