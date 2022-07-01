package account

import (
	"game-blockchain-server/serializer"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"os"
)

type UpdatePasswordService struct {
	Ks          *keystore.KeyStore
	OldPassWord string
	NewPassWord string
}

func (service *UpdatePasswordService) UpdateInitialPassword() serializer.Response {

	service.Ks = keystore.NewKeyStore(os.Getenv("KEY_DIR"), keystore.StandardScryptN, keystore.StandardScryptP)
	err := service.Ks.Update(service.Ks.Accounts()[0], "Creation password", service.NewPassWord)
	if err != nil {
		return serializer.Response{
			Code:  400,
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
	}
}

func (service *UpdatePasswordService) UpdatePassword() serializer.Response {
	service.Ks = keystore.NewKeyStore(os.Getenv("KEY_DIR"), keystore.StandardScryptN, keystore.StandardScryptP)
	err := service.Ks.Update(service.Ks.Accounts()[0], service.OldPassWord, service.NewPassWord)
	if err != nil {
		return serializer.Response{
			Code:  400,
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
	}
}
