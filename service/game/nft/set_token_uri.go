package nft

import (
	"game-blockchain-server/constants"
	"game-blockchain-server/serializer"
	"game-blockchain-server/service/signature"
	tx "game-blockchain-server/service/transaction"
	"game-blockchain-server/utils"
	"os"
)

type SetTokenURI struct {
	TokenID        string `form:"token_id" binding:"required"`
	TokenURI       string `form:"token_uri" binding:"required"`
	ContractNumber int    `form:"contract_number" binding:"required"`
}

func (service *SetTokenURI) SetTokenURI() serializer.Response {
	methodID := utils.GetTxMethodName("setTokenURl(uint256,string)")

	tokenID := utils.GetTxUint256(service.TokenID)
	tokenURI := utils.GetTxString(service.TokenURI)

	var data []byte
	data = append(data, methodID...)
	data = append(data, tokenID...)
	data = append(data, tokenURI...)

	contractAddress, err := constants.GetContractAddress(service.ContractNumber)

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
