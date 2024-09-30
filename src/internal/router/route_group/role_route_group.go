package route_group

import (
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewRoleRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/role")

	r.POST("", handler.RoleHandler.Create)

	r.GET("", handler.RoleHandler.FindAllRole)
	r.GET(":roleId", handler.RoleHandler.FindRoleById)

	r.PUT(":roleId", handler.RoleHandler.Update)

	r.DELETE(":roleId", handler.RoleHandler.Remove)
}