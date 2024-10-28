package middlewares

import (
	"net/http"
	"src/internal/security"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	jwtSecurity := security.NewJwtTokenService()
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			context.Abort()
			return
		}

		tokenStr := authHeader[len("Bearer "):]

		token, err := jwtSecurity.ValidateToken(tokenStr)

		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err})
			context.Abort()
			return
		}
		context.Set("userId", token.UserID)
		context.Next()
	}
}
