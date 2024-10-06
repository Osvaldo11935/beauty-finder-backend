package route_group

import (
	"src/internal/delivery/http/middlewares"
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewAddressRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/address")

	r.POST("/user/:userId", handler.AddressHandler.Create)

	r.GET("/user/:userId", middlewares.Auth(), handler.AddressHandler.FindAddressByUserId)

	r.PUT("/user/:userId", middlewares.Auth(), handler.AddressHandler.Update)

	r.DELETE("/user/:userId", middlewares.Auth(), handler.AddressHandler.Remove)
}
