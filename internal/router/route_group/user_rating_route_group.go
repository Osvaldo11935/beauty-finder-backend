package route_group

import (
	"src/internal/delivery/http/middlewares"
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewUserRatingRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/user-rating")

	r.POST("", middlewares.Auth(), handler.UserRatingHandler.Create)
	r.POST("userEvaluator/:userEvaluatorId/dispatch-assessment-notification", middlewares.Auth(), handler.UserRatingHandler.DispatchAssessmentNotification)

	r.GET("", middlewares.Auth(), handler.UserRatingHandler.FindAllUserRating)
	r.GET(":userRatingId", middlewares.Auth(), handler.UserRatingHandler.FindUserRatingById)
	r.GET("user/:userId", middlewares.Auth(), handler.UserRatingHandler.FindUserRatingByUserId)

	r.PUT(":userRatingId", middlewares.Auth(), handler.UserRatingHandler.Update)

	r.DELETE(":userRatingId", middlewares.Auth(), handler.UserRatingHandler.Remove)
}
