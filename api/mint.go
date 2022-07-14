package api

import (
	service "game-blockchain-server/service/game/mint"
	"github.com/gin-gonic/gin"
)

func MintERC20(c *gin.Context) {
	var service service.MintERC20Service
	if err := c.ShouldBind(&service); err == nil {
		res := service.MintERC20()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func MintERC721(c *gin.Context) {
	var service service.MintERC721Service
	if err := c.ShouldBind(&service); err == nil {
		res := service.MintSoul()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
