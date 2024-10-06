package route_group

import (
	"src/internal/delivery/http/middlewares"
	"src/internal/setup"
	"github.com/gin-gonic/gin"
)

func NewPersonRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/person")

	r.POST("/user/:userId", middlewares.Auth(), handler.PersonHandler.Create)

	r.GET("/user/:userId", middlewares.Auth(), handler.PersonHandler.FindPersonByUserId)
	r.GET("national-registry/:nationalRegistry", handler.PersonHandler.FindPersonByNationalRegistry)

	r.PUT("/user/:userId", middlewares.Auth(), handler.PersonHandler.Update)

	r.DELETE("/user/:userId", middlewares.Auth(), handler.PersonHandler.Remove)
}