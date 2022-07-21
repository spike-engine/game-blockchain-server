package withdrawal

import (
	"game-blockchain-server/constants"
	"game-blockchain-server/serializer"
	"game-blockchain-server/service/signature"
	tx "game-blockchain-server/service/transaction"
	"game-blockchain-server/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"os"
)

type WithdrawalERC721Service struct {
	FromAddress    string `form:"from_address" binding:"required"`
	ToAddress      string `form:"to_address" binding:"required"`
	TokenID        string `form:"token_id" binding:"required"`
	ContractNumber int    `form:"contract_number" binding:"required"`
}

func (service *WithdrawalERC721Service) WithdrawalSoul() serializer.Response {
	methodID := utils.GetTxMethodName("transferFrom(address,address,uint256)")

	fromAddress := utils.GetTxAddress(service.FromAddress)

	toAddress := utils.GetTxAddress(service.ToAddress)

	paddedAmount := utils.GetTxUint256(service.TokenID)

	var data []byte
	data = append(data, methodID...)
	data = append(data, fromAddress...)
	data = append(data, toAddress...)
	data = append(data, paddedAmount...)

	log.Info("====Spike log: ", "methodID: ", hexutil.Encode(methodID), "from_address: ", hexutil.Encode(fromAddress), "to_address: ", hexutil.Encode(toAddress), "amount: ", hexutil.Encode(paddedAmount))

	contractAddress, err := constants.GetContractAddress(service.ContractNumber)

	spikeTx := &utils.SpikeTx{
		Data: data,
		To:   contractAddress,
	}
	transaction, err := spikeTx.ConstructionTransaction()
	if err != nil {
		log.Error("====Spike log: ", err)
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
		log.Error("====Spike log: ", err)
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
		log.Error("====Spike log: ", err)
		return serializer.Response{
			Code:  405,
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: signedTx.Hash(),
	}
}
