package signature

import (
	"context"
	"game-blockchain-server/config"
	"game-blockchain-server/serializer"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"os"
)

type SignTxService struct {
	Ks       *keystore.KeyStore
	PassWord string
	TX       []byte
}

func (service *SignTxService) SignSeparateTX() serializer.Response {
	service.Ks = keystore.NewKeyStore(os.Getenv("KEY_DIR"), keystore.StandardScryptN, keystore.StandardScryptP)
	unsignTx := new(types.Transaction)

	err := unsignTx.UnmarshalBinary(service.TX)
	if err != nil {
		return serializer.Response{
			Code:  401,
			Error: err.Error(),
		}
	}

	client, err := config.GetUsefulBscNode()
	if err != nil {
		return serializer.Response{
			Code:  402,
			Error: err.Error(),
		}
	}

	chainid, err := client.ChainID(context.Background())
	if err != nil {
		return serializer.Response{
			Code:  402,
			Error: err.Error(),
		}
	}

	signedTx, err := service.Ks.SignTxWithPassphrase(service.Ks.Accounts()[0], service.PassWord, unsignTx, chainid)
	if err != nil {
		return serializer.Response{
			Code:  400,
			Error: err.Error(),
		}
	}

	signedTxByte, err := signedTx.MarshalBinary()
	if err != nil {
		return serializer.Response{
			Code:  402,
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Data: signedTxByte,
	}
}
