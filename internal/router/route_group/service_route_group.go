package route_group

import (
	"src/internal/delivery/http/middlewares"
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewServiceRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/service")

	r.POST("", middlewares.Auth(), handler.ServiceHandler.Create)

	r.GET("", middlewares.Auth(), handler.ServiceHandler.FindAllService)
	r.GET(":serviceId", middlewares.Auth(), handler.ServiceHandler.FindServiceById)
	r.GET("category/:categoryId", middlewares.Auth(), handler.ServiceHandler.FindServiceByCategoryId)
	r.GET("provider/:providerId", middlewares.Auth(), handler.ServiceHandler.FindServiceByProviderId)

	r.PUT(":serviceId", middlewares.Auth(), handler.ServiceHandler.Update)

	r.DELETE(":serviceId", middlewares.Auth(), handler.ServiceHandler.Remove)
}