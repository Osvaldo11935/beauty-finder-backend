package route_group

import (
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewCategoryRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/category")

	r.POST("", handler.ServiceCategoryHandler.Create)

	r.GET("", handler.ServiceCategoryHandler.FindAllServiceCategory)
	r.GET(":categoryId", handler.ServiceCategoryHandler.FindServiceCategoryById)

	r.PUT(":categoryId", handler.ServiceCategoryHandler.Update)

	r.DELETE(":categoryId", handler.ServiceCategoryHandler.Remove)
}