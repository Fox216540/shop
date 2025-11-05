package middleware

import (
	"github.com/Fox216540/shop/apigateway-service/domain/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type JWTMiddleware struct {
	authClient auth.Client
}

func (jwt *JWTMiddleware) DecodeAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		authHeader := c.Request.Header.Get("Authorization")
		log.Println(authHeader)
		if strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}
		log.Println(token)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized 1"})
			c.Abort()
			return
		}
		id, err := jwt.authClient.DecodeAccessTokenOfUser(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("user_id", id)
		c.Next()
	}
}
