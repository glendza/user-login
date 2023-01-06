package controllers

import (
	"net/http"

	"glendza/login-code-challenge/models"
	"glendza/login-code-challenge/requests"
	"glendza/login-code-challenge/responses"
	"glendza/login-code-challenge/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (u AuthController) Login(c *gin.Context) {
	data := requests.AuthRequest{}

	if err := c.BindJSON(&data); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := models.AuthenticateUser(data.Username, data.Password)

	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	token, _ := services.GenerateJWT(user)
	response := responses.AuthResponse{
		Token: token,
	}
	c.IndentedJSON(http.StatusOK, response)
}
