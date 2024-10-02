package route_group

import (
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewAddressRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/address")

	r.POST("/user/:userId", handler.AddressHandler.Create)

	r.GET("/user/:userId", handler.AddressHandler.FindAddressByUserId)

	r.PUT("/user/:userId", handler.AddressHandler.Update)

	r.DELETE("/user/:userId", handler.AddressHandler.Remove)
}