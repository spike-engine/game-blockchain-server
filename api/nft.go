package api

import (
	service "game-blockchain-server/service/game/nft"
	"github.com/gin-gonic/gin"
)

func SetBaseTokenURI(c *gin.Context) {
	var service service.SetBaseTokenURI
	if err := c.ShouldBind(&service); err == nil {
		res := service.SetBaseTokenURI()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func SetTokenURI(c *gin.Context) {
	var service service.SetTokenURI
	if err := c.ShouldBind(&service); err == nil {
		res := service.SetTokenURI()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
