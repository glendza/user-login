package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) RetrieveSelf(c *gin.Context) {
	// A slight mixup of domain and presentation layer, again, for the demo purposes
	user, _ := c.Get("x-user")
	c.IndentedJSON(http.StatusOK, user)
}
