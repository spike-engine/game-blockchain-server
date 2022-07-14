package withdrawal

import (
	"game-blockchain-server/constants"
	"game-blockchain-server/serializer"
	"game-blockchain-server/service/signature"
	tx "game-blockchain-server/service/transaction"
	"game-blockchain-server/utils"
	"os"
)

type WithdrawalERC20Service struct {
	ToAddress      string `form:"to_address",json:"to_address"`
	Amount         string `form:"amount",json:"amount"`
	ContractNumber int    `form:"contract_number",json:"contract_number"`
}

func (service *WithdrawalERC20Service) WithdrawalERC20() serializer.Response {
	methodID := utils.GetTxMethodName("transfer(address,uint256)")

	paddedAddress := utils.GetTxAddress(service.ToAddress)

	paddedAmount := utils.GetTxAmount(service.Amount)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	contractAddress, err := constants.GetContractAddress(service.ContractNumber)

	if err != nil {
		return serializer.Response{
			Code:  401,
			Error: err.Error(),
		}
	}

	spikeTx := &utils.SpikeTx{
		Data: data,
		To:   contractAddress,
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
		Data: signedTx.Hash(),
	}

}
