package middlewares

import (
	"glendza/login-code-challenge/models"
	"glendza/login-code-challenge/services"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Token")
		userId, err := services.ParseJWT(token)

		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		user := models.GetUserById(userId)
		c.Set("x-user", *user)

		c.Next()
	}
}
