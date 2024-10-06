package middlewares

import (
	"net/http"
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	"src/internal/delivery/http/validators"
	"src/internal/domain/errors"
	"github.com/gin-gonic/gin"
)

func ValidateUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		var request models_requests_posts.CreateUserRequest
    
		if err := c.ShouldBindJSON(&request); err != nil {
			customErr := errors.UnknownCreateUserError(err.Error())
			c.JSON(http.StatusBadRequest, customErr)
			c.Abort()
			return
		}
        
		validationErrors := validators.ValidateCreateUser(&request)
		if validationErrors != nil {
			c.JSON(http.StatusBadRequest, validationErrors)
			c.Abort()
			return
		}
		c.Set("request", request)
		c.Next()
	}
}
