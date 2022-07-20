package utils

import (
	"context"
	"game-blockchain-server/config"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"os"
)

type SpikeTx struct {
	Data []byte
	To   string
}

func (service *SpikeTx) ConstructionTransaction() (*types.Transaction, error) {

	client, err := config.GetUsefulBscNode()
	if err != nil {
		return nil, err

	}

	keyStore := keystore.NewKeyStore(os.Getenv("KEY_DIR"), keystore.StandardScryptN, keystore.StandardScryptP)

	nonce, err := client.PendingNonceAt(context.Background(), keyStore.Accounts()[0].Address)
	if err != nil {
		return nil, err
	}

	toAddress := common.HexToAddress(service.To)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: keyStore.Accounts()[0].Address,
		To:   &toAddress,
		Data: service.Data,
	})
	if err != nil {
		return nil, err
	}

	tx := types.NewTx(
		&types.LegacyTx{
			Nonce:    nonce,
			Gas:      gasLimit,
			GasPrice: gasPrice,
			Data:     service.Data,
			To:       &toAddress,
		})
	return tx, nil
}
