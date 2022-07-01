package main

import (
	"game-blockchain-server/config"
	"game-blockchain-server/server"
)

func main() {
	config.Init()
	r := server.NewRouter()
	r.Run(":3000")
}
