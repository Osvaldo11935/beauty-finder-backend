package route_group

import (
	"src/internal/delivery/http/middlewares"
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewAttachmentTypeRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/attachment-type")

	r.POST("", middlewares.Auth(), handler.AttachmentTypeHandler.Create)

	r.GET("", middlewares.Auth(), handler.AttachmentTypeHandler.FindAllAttachmentType)
	r.GET(":attachmentTypeId", middlewares.Auth(), handler.AttachmentTypeHandler.FindAttachmentTypeById)
	

	r.PUT(":attachmentTypeId", middlewares.Auth(), handler.AttachmentTypeHandler.Update)

	r.DELETE(":attachmentTypeId", middlewares.Auth(), handler.AttachmentTypeHandler.Remove)
}