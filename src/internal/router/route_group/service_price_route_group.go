package route_group

import (
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewServicePriceRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/price")

	r.POST("", handler.ServicePriceHandler.Create)

	r.GET(":priceId", handler.ServicePriceHandler.FindServicePriceById)
	r.GET("service/:serviceId", handler.ServicePriceHandler.FindServicePriceByServiceId)
	
	r.PUT(":priceId", handler.ServicePriceHandler.Update)

	r.DELETE(":priceId", handler.ServicePriceHandler.Remove)
}