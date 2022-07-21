package nft

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

var log = logger.Logger("nft")

type SetBaseTokenURI struct {
	BaseURI        string `form:"base_uri" binding:"required"`
	ContractNumber int    `form:"contract_number" binding:"required"`
}

func (service *SetBaseTokenURI) SetBaseTokenURI() serializer.Response {
	methodID := utils.GetTxMethodName("setBaseTokenURI(string)")

	baseURI := utils.GetTxString(service.BaseURI)

	offset := utils.GetOffset(1)

	var data []byte
	data = append(data, methodID...)
	data = append(data, offset...)
	data = append(data, baseURI...)

	log.Info("====Spike log: ", "methodID: ", hexutil.Encode(methodID), "baseURI: ", hexutil.Encode(baseURI), "offset: ", offset)

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
