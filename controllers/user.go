package controllers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/awesome-life/god/models"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) Retrieve(c *gin.Context) {
}
