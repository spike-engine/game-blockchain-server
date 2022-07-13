package transaction

import (
	"context"
	"game-blockchain-server/config"
	"github.com/ethereum/go-ethereum/core/types"
)

type BroadcastService struct {
	SignedTX *types.Transaction
}

func (service *BroadcastService) SendTransaction() error {

	client, err := config.GetUsefulBscNode()
	if err != nil {
		return err
	}

	err = client.SendTransaction(context.Background(), service.SignedTX)
	if err != nil {
		return err
	}

	return nil
}
