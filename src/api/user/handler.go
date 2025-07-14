package userhandler

import "github.com/gin-gonic/gin"

func UserHandler(r *gin.Engine) {
	// Здесь вы можете определить обработчики для маршрутов, связанных с пользователями
	// Например:
	r.GET("/users", getUsers)
	r.POST("/users", createUser)

	r.DELETE("/users/:id", deleteUser)

	// Пример простого обработчика
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "List of users",
		})
	})

}
