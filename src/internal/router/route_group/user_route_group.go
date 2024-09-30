package route_group

import (
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewUserRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/user")

	r.POST("admin", handler.UserHandler.CreateAdmin)
	r.POST("client", handler.UserHandler.CreateClient)
	r.POST("service-provider", handler.UserHandler.CreateServiceProvider)
	r.POST(":userId/service-provided", handler.UserHandler.CreateServiceProvided)

	r.GET("", handler.UserHandler.FindAllUser)
	r.GET(":useId", handler.UserHandler.FindUserById)
	r.GET("service/:serviceId", handler.UserHandler.FindUserByServiceId)

	r.PUT(":UserId", handler.UserHandler.Update)

	r.DELETE(":UserId", handler.UserHandler.Remove)
}
