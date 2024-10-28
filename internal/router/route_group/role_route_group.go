package route_group

import (
	"src/internal/delivery/http/middlewares"
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewRoleRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/role")

	r.POST("", middlewares.Auth(), handler.RoleHandler.Create)

	r.GET("", middlewares.Auth(), handler.RoleHandler.FindAllRole)
	r.GET(":roleId", middlewares.Auth(), handler.RoleHandler.FindRoleById)

	r.PUT(":roleId", middlewares.Auth(), handler.RoleHandler.Update)

	r.DELETE(":roleId", middlewares.Auth(), handler.RoleHandler.Remove)
}