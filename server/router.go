package server

import (
	"game-blockchain-server/api"
	"game-blockchain-server/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.LoggerToFile())

	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "pong")
		})

		account := v1.Group("/account")
		{
			account.GET("update/initpassword", api.UpdateInitialPassword)
			account.GET("update/password", api.UpdatePassword)
		}

		ipfs := v1.Group("/signature")
		{
			ipfs.POST("pin/file", api.SignSeparateTX)
		}
	}
	return r
}
