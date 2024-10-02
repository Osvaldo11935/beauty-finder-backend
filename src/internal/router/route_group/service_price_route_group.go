package route_group

import (
	"src/internal/delivery/http/middlewares"
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewServicePriceRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/price")

	r.POST("", middlewares.Auth(), handler.ServicePriceHandler.Create)

	r.GET(":priceId", middlewares.Auth(), handler.ServicePriceHandler.FindServicePriceById)
	r.GET("service/:serviceId", middlewares.Auth(), handler.ServicePriceHandler.FindServicePriceByServiceId)
	
	r.PUT(":priceId", middlewares.Auth(), handler.ServicePriceHandler.Update)

	r.DELETE(":priceId", middlewares.Auth(), handler.ServicePriceHandler.Remove)
}