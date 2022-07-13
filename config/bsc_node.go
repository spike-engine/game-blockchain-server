package config

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"strings"
)

func GetUsefulBscNode() (*ethclient.Client, error) {
	var err error
	//nodes := strings.Split(os.Getenv("BSC_NODE"), ",")
	nodes := strings.Split("https://bsctestapi.terminet.io/rpc,https://data-seed-prebsc-1-s2.binance.org:8545", ",")
	for _, node := range nodes {
		client, err := ethclient.Dial(node)
		if err != nil {
			continue
		}
		return client, nil
	}
	return nil, err
}
