package account

import (
	"game-blockchain-server/serializer"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"os"
)

type UpdatePasswordService struct {
	OldPassWord string `form:"old_password" json:"old_password"`
	NewPassWord string `form:"new_password" json:"new_password"`
}

func (service *UpdatePasswordService) UpdateInitialPassword() serializer.Response {

	ks := keystore.NewKeyStore(os.Getenv("KEY_DIR"), keystore.StandardScryptN, keystore.StandardScryptP)
	err := ks.Update(ks.Accounts()[0], "Creation password", service.NewPassWord)
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
	ks := keystore.NewKeyStore(os.Getenv("KEY_DIR"), keystore.StandardScryptN, keystore.StandardScryptP)
	err := ks.Update(ks.Accounts()[0], service.OldPassWord, service.NewPassWord)
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
