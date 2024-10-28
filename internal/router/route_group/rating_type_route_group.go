package route_group

import (
	"src/internal/delivery/http/middlewares"
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewRatingTypeRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/rating-type")

	r.POST("", middlewares.Auth(), handler.RatingTypeHandler.Create)

	r.GET("", middlewares.Auth(), handler.RatingTypeHandler.FindAllRatingType)
	r.GET(":ratingTypeId", middlewares.Auth(), handler.RatingTypeHandler.FindRatingTypeById)

	r.PUT(":ratingTypeId", middlewares.Auth(), handler.RatingTypeHandler.Update)

	r.DELETE(":ratingTypeId", middlewares.Auth(), handler.RatingTypeHandler.Remove)
}