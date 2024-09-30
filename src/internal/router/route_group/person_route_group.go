package route_group

import (
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewPersonRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/person")

	r.POST("/user/:userId", handler.PersonHandler.Create)

	r.GET("/user/:userId", handler.PersonHandler.FindPersonByUserId)

	r.PUT("/user/:userId", handler.PersonHandler.Update)

	r.DELETE("/user/:userId", handler.PersonHandler.Remove)
}