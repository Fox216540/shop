package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"shop/src/domain/jwt"
	"strings"
)

func JWTMiddleware(jwt jwt.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		authHeader := c.Request.Header.Get("Authorization")
		log.Println(authHeader)
		if strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}
		log.Println(token)
		if token == "" {
			c.JSON(401, gin.H{"error": "Unauthorized 1"})
			c.Abort()
			return
		}
		u, err := jwt.DecodeAccessToken(token)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("user_id", u.UserID)
		c.Next()
	}
}
