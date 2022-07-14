package config

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
	"strings"
)

func GetUsefulBscNode() (*ethclient.Client, error) {
	var err error
	nodes := strings.Split(os.Getenv("BSC_NODE"), ",")
	//nodes := strings.Split("https://data-seed-prebsc-1-s2.binance.org:8545,https://data-seed-prebsc-1-s1.binance.org:8545", ",")
	for _, node := range nodes {
		client, err := ethclient.Dial(node)
		if err != nil {
			continue
		}
		return client, nil
	}
	return nil, err
}
