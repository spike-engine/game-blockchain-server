package api

import (
	service "game-blockchain-server/service/account"
	"github.com/gin-gonic/gin"
)

func UpdateInitialPassword(c *gin.Context) {

	var service service.UpdatePasswordService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateInitialPassword()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func UpdatePassword(c *gin.Context) {
	var service service.UpdatePasswordService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdatePassword()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
