package route_group

import (
	"src/internal/delivery/http/middlewares"
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewCompanyRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/company")

	r.POST("", handler.CompanyHandler.Create)

	r.GET("", middlewares.Auth(), handler.CompanyHandler.FindAllCompany)
	r.GET(":companyId", middlewares.Auth(), handler.CompanyHandler.FindCompanyById)
	r.GET("national-registry/:nationalRegistry", handler.CompanyHandler.FindCompanyByNationalRegistry)

	r.PUT(":companyId", middlewares.Auth(), handler.CompanyHandler.Update)
	r.DELETE(":companyId", middlewares.Auth(), handler.CompanyHandler.Remove)
}
