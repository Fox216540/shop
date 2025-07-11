package user_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	dto "shop/src/api/dto/user_dto"
)

func UserHandler(r *gin.Engine) {
	req := dto.UserLoginRequest{}
	req.Password = "123456"
	req.UsernameOrEmail = "email"
	r.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, req)
	})
}
