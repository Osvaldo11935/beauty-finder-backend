package route_group

import (
	"src/internal/setup"
	"github.com/gin-gonic/gin"
)

func NewServiceRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/service")

	r.POST("", handler.ServiceHandler.Create)

	r.GET("", handler.ServiceHandler.FindAllService)
	r.GET(":serviceId", handler.ServiceHandler.FindServiceById)
	r.GET("category/:categoryId", handler.ServiceHandler.FindServiceByCategoryId)
	r.GET("provider/:providerId", handler.ServiceHandler.FindServiceByProviderId)

	r.PUT(":serviceId", handler.ServiceHandler.Update)

	r.DELETE(":serviceId", handler.ServiceHandler.Remove)
}