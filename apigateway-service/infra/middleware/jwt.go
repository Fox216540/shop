package middleware

import (
	"github.com/Fox216540/shop/apigateway-service/app/tokenDecoder"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const ContextUserID = "user_id"

type JWTMiddleware struct {
	tokenDecoder tokenDecoder.UseCase
}

func (jwt *JWTMiddleware) unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, NewMiddlewareError(nil))
	c.Abort()
}

func (jwt *JWTMiddleware) DecodeAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			jwt.unauthorized(c)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		id, err := jwt.tokenDecoder.DecodeAccessToken(token)
		if err != nil {
			jwt.unauthorized(c)
			return
		}
		c.Set(ContextUserID, id)
		c.Next()
	}
}
