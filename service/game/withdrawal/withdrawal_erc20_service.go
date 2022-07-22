package withdrawal

import (
	"game-blockchain-server/constants"
	"game-blockchain-server/serializer"
	"game-blockchain-server/service/signature"
	tx "game-blockchain-server/service/transaction"
	"game-blockchain-server/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	logger "github.com/ipfs/go-log"
	"os"
)

var log = logger.Logger("withdrawal")

type WithdrawalERC20Service struct {
	ToAddress      string `form:"to_address" json:"to_address" binding:"required"`
	Amount         string `form:"amount" json:"amount" binding:"required"`
	ContractNumber int    `form:"contract_number" json:"contract_number" binding:"required"`
}

func (service *WithdrawalERC20Service) WithdrawalERC20() serializer.Response {
	methodID := utils.GetTxMethodName("transfer(address,uint256)")

	paddedAddress := utils.GetTxAddress(service.ToAddress)

	paddedAmount := utils.GetTxUint256(service.Amount)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	log.Info("====Spike log: ", "methodID: ", hexutil.Encode(methodID), "address: ", hexutil.Encode(paddedAddress), "amount: ", hexutil.Encode(paddedAmount))

	contractAddress, err := constants.GetContractAddress(service.ContractNumber)

	if err != nil {
		log.Error("====Spike log: ", err)
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
