package withdrawal

import (
	"game-blockchain-server/constants"
	"game-blockchain-server/serializer"
	"game-blockchain-server/service/signature"
	tx "game-blockchain-server/service/transaction"
	"game-blockchain-server/utils"
	"os"
)

type WithdrawalERC721Service struct {
	FromAddress string
	ToAddress   string
	TokenID     string
}

func (service *WithdrawalERC721Service) WithdrawalSoul() serializer.Response {
	methodID := utils.GetTxMethodName("transfer(address,uint256)")

	fromAddress := utils.GetTxAddress(service.FromAddress)

	toAddress := utils.GetTxAddress(service.ToAddress)

	paddedAmount := utils.GetTxAmount(service.TokenID)

	var data []byte
	data = append(data, methodID...)
	data = append(data, fromAddress...)
	data = append(data, toAddress...)
	data = append(data, paddedAmount...)

	spikeTx := &utils.SpikeTx{
		Data: data,
		To:   constants.SOUL_CONTRACT_ADDRESS_TESTNET,
	}
	transaction, err := spikeTx.ConstructionTransaction()
	if err != nil {
		return serializer.Response{
			Code:  402,
			Error: err.Error(),
		}
	}

	SignTxService := &signature.SignTxService{
		PassWord: os.Getenv("OWNER_PW"),
		TX:       transaction,
	}

	signedTx, err := SignTxService.SignSeparateTX()
	if err != nil {
		return serializer.Response{
			Code:  403,
			Error: err.Error(),
		}
	}

	Broad := &tx.BroadcastService{
		SignedTX: signedTx,
	}

	err = Broad.SendTransaction()
	if err != nil {
		return serializer.Response{
			Code: 405,
		}
	}

	return serializer.Response{
		Code: 200,
	}
}
