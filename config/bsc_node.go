package config

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
	"strings"
)

func GetUsefulBscNode() (*ethclient.Client, error) {
	var err error
	nodes := strings.Split(os.Getenv("BSC_NODE"), ",")
	for _, node := range nodes {
		client, err := ethclient.Dial(node)
		if err != nil {
			continue
		}
		return client, nil
	}
	return nil, err
}
