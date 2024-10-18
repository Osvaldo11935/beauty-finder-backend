package route_group

import (
	"src/internal/delivery/websocket/handlers"
	"src/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewChatRouteGroup(route *gin.RouterGroup, setup *usecase.UserUseCase, pool *usecase.Pool) {
	r := route.Group("/ws")

	r.GET("client", func(ctx *gin.Context) {
		handlers.WebSocketHandler(pool, setup, ctx.Writer, ctx.Request)
	})
	r.GET("service-provider", func(ctx *gin.Context) {
		handlers.WebSocketHandler(pool, setup, ctx.Writer, ctx.Request)
	})
}
