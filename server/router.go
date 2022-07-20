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
			account.POST("update/initpassword", api.UpdateInitialPassword)
			account.POST("update/password", api.UpdatePassword)
		}

		game := v1.Group("/game")
		{
			game.POST("mint/erc20", api.MintERC20)
			game.POST("mint/erc721", api.MintERC721)
			game.POST("withdrawal/erc20", api.WithdrawalERC20)
			game.POST("withdrawal/erc721", api.WithdrawalERC721)
			game.POST("nft/baseuri", api.SetBaseTokenURI)
			game.POST("nft/tokenuri", api.SetTokenURI)

		}
	}
	return r
}
