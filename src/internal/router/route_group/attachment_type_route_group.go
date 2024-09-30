package route_group

import (
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewAttachmentTypeRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/attachment-type")

	r.POST("", handler.AttachmentTypeHandler.Create)

	r.GET("", handler.AttachmentTypeHandler.FindAllAttachmentType)
	r.GET(":attachmentTypeId", handler.AttachmentTypeHandler.FindAttachmentTypeById)
	

	r.PUT(":attachmentTypeId", handler.AttachmentTypeHandler.Update)

	r.DELETE(":attachmentTypeId", handler.AttachmentTypeHandler.Remove)
}