package signature

import (
	"context"
	"game-blockchain-server/config"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"os"
)

type SignTxService struct {
	PassWord string
	TX       *types.Transaction
}

func (service *SignTxService) SignSeparateTX() (*types.Transaction, error) {
	ks := keystore.NewKeyStore(os.Getenv("KEY_DIR"), keystore.StandardScryptN, keystore.StandardScryptP)

	client, err := config.GetUsefulBscNode()
	if err != nil {
		return nil, err
	}

	chainid, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	signedTx, err := ks.SignTxWithPassphrase(ks.Accounts()[0], service.PassWord, service.TX, chainid)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}
