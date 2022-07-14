package api

import (
	service "game-blockchain-server/service/game/withdrawal"
	"github.com/gin-gonic/gin"
)

func WithdrawalERC20(c *gin.Context) {
	var service service.WithdrawalERC20Service
	if err := c.ShouldBind(&service); err == nil {
		res := service.WithdrawalERC20()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func WithdrawalERC721(c *gin.Context) {
	var service service.WithdrawalERC721Service
	if err := c.ShouldBind(&service); err == nil {
		res := service.WithdrawalSoul()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
