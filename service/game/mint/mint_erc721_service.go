package mint

import (
	"fmt"
	"game-blockchain-server/constants"
	"game-blockchain-server/serializer"
	"game-blockchain-server/service/signature"
	tx "game-blockchain-server/service/transaction"
	"game-blockchain-server/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"os"
)

type MintERC721Service struct {
	TokenID string `form:"token_id" binding:"required"`
	//  toAddress is nft owner, Can be the owner or administrator of the contract
	ToAddress string `form:"to_address" binding:"required"`

	TokenURI       string `form:"token_uri" binding:"required"`
	ContractNumber int    `form:"contract_number" binding:"required"`
}

func (service *MintERC721Service) MintSoul() serializer.Response {
	methodID := utils.GetTxMethodName("mint(uint256,address,string)")

	paddedAddress := utils.GetTxAddress(service.ToAddress)

	paddedAmount := utils.GetTxUint256(service.TokenID)

	tokenURI := utils.GetTxString(service.TokenURI)

	offset := utils.GetOffset(3)

	fmt.Println("methodID: ", hexutil.Encode(methodID), "address: ", hexutil.Encode(paddedAddress), "amount: ", hexutil.Encode(paddedAmount), "tokenURL: ", hexutil.Encode(tokenURI), "offset: ", offset)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAmount...)
	data = append(data, paddedAddress...)
	data = append(data, offset...)
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
			Code:  405,
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: signedTx.Hash(),
	}
}
