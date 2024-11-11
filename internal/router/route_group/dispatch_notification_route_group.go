package route_group

import (
	"src/internal/delivery/http/middlewares"
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewDispatchNotificationRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("")

	r.POST("appointment/:appointmentId/service/:serviceId/dispatch-notification", middlewares.Auth(), handler.DispatchNotificationHandler.DispatchServiceNotification)
	r.POST("user-rating/userEvaluator/:userEvaluatorId/dispatch-assessment-notification", middlewares.Auth(), handler.DispatchNotificationHandler.DispatchAssessmentNotification)
	r.POST("userApplicant/:userApplicantId/userRequired/:userRequiredId/dispatch-start-conversation-notification", middlewares.Auth(), handler.DispatchNotificationHandler.DispatchStartConversationNotification)

}

