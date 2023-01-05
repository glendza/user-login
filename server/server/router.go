package server

import (
	"glendza/login-code-challenge/controllers"
	"glendza/login-code-challenge/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("api/v1")
	{
		authGroup := v1.Group("auth")
		{
			auth := new(controllers.AuthController)
			authGroup.GET("/login", auth.Login)
		}

		userGroup := v1.Group("user").Use(middlewares.AuthMiddleware())
		{
			user := new(controllers.UserController)
			userGroup.GET("/self", user.RetrieveSelf)
		}
	}
	return router
}
