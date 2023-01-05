package middlewares

import (
	"glendza/login-code-challenge/models"
	"glendza/login-code-challenge/services"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Token")
		token, err := services.ParseJWT(tokenString)

		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		userId := services.ExtractUserId(token)
		user := models.GetUserById(userId)
		c.Set("x-user", *user)

		c.Next()
	}
}
