package route_group

import (
	"src/internal/delivery/http/middlewares"
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewCategoryRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/category")

	r.POST("", middlewares.Auth(), handler.ServiceCategoryHandler.Create)

	r.GET("", middlewares.Auth(), handler.ServiceCategoryHandler.FindAllServiceCategory)
	r.GET(":categoryId", middlewares.Auth(), handler.ServiceCategoryHandler.FindServiceCategoryById)

	r.PUT(":categoryId", middlewares.Auth(), handler.ServiceCategoryHandler.Update)

	r.DELETE(":categoryId", middlewares.Auth(), handler.ServiceCategoryHandler.Remove)
}