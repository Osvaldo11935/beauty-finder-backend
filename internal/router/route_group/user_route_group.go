package route_group

import (
	"src/internal/delivery/http/middlewares"
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewUserRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/user")

	r.POST("token", handler.UserHandler.FindToken)
	r.POST("admin", middlewares.Auth(), middlewares.ValidateUserMiddleware(), handler.UserHandler.CreateAdmin)
	r.POST("client", middlewares.ValidateUserMiddleware(), handler.UserHandler.CreateClient)
	r.POST("service-provider", middlewares.ValidateUserMiddleware(), handler.UserHandler.CreateServiceProvider)
	r.POST(":userId/service-provided", middlewares.Auth(), handler.UserHandler.CreateServiceProvided)
	r.POST(":userId/fcm-token", middlewares.Auth(), handler.UserHandler.CreateFcmToken)

	r.GET("", middlewares.Auth(), handler.UserHandler.FindAllUser)
	r.GET("user-info", middlewares.Auth(), handler.UserHandler.FindUserById)
	r.GET("service/:serviceId", middlewares.Auth(), handler.UserHandler.FindUserByServiceId)
	r.GET("service/:serviceId/lat/:latitude/long/:longitude", middlewares.Auth(), handler.UserHandler.FindUsersNearBy)
	r.PUT(":UserId", middlewares.Auth(), handler.UserHandler.Update)

	r.DELETE(":UserId", middlewares.Auth(), handler.UserHandler.Remove)
}
