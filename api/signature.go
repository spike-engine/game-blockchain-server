package api

import (
	service "game-blockchain-server/service/signature"
	"github.com/gin-gonic/gin"
)

func SignSeparateTX(c *gin.Context) {
	var service service.SignTxService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SignSeparateTX()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
