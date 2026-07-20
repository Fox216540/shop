package middleware

import (
	shopApiGen "github.com/Fox216540/shop/api/gen"
	"github.com/Fox216540/shop/apigateway-service/app/tokenDecoder"
	"github.com/Fox216540/shop/apigateway-service/core/transport"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const ContextUserID = transport.UserIDKey

type JWTMiddleware struct {
	tokenDecoder tokenDecoder.UseCase
}

func NewJWTMiddleware(tokenDecoder tokenDecoder.UseCase) *JWTMiddleware {
	return &JWTMiddleware{
		tokenDecoder: tokenDecoder,
	}
}

func (jwt *JWTMiddleware) unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, NewMiddlewareError(nil))
	c.Abort()
}

func (jwt *JWTMiddleware) Security() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Get(shopApiGen.BearerAuthScopes); !ok {
			c.Next()
			return
		}

		jwt.decodeAccessToken(c)
	}
}

func (jwt *JWTMiddleware) DecodeAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt.decodeAccessToken(c)
	}
}

func (jwt *JWTMiddleware) decodeAccessToken(c *gin.Context) {
	authHeader := c.GetHeader(transport.AuthorizationHeader)
	if !strings.HasPrefix(authHeader, transport.BearerPrefix) {
		jwt.unauthorized(c)
		return
	}

	token := strings.TrimPrefix(authHeader, transport.BearerPrefix)
	id, err := jwt.tokenDecoder.DecodeAccessToken(token)
	if err != nil {
		jwt.unauthorized(c)
		return
	}

	c.Set(ContextUserID, id)
	c.Next()
}
