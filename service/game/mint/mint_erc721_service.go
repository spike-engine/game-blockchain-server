package mint

import (
	"fmt"
	"game-blockchain-server/constants"
	"game-blockchain-server/serializer"
	"game-blockchain-server/service/signature"
	tx "game-blockchain-server/service/transaction"
	"game-blockchain-server/utils"
	"os"
)

type MintERC721Service struct {
	TokenID string `form:"token_id" binding:"required"`
	//  toAddress is nft owner, Can be the owner or administrator of the contract
	ToAddress string `form:"to_address" binding:"required"`
}

func (service *MintERC721Service) MintSoul() serializer.Response {
	methodID := utils.GetTxMethodName("mint(uint256,address)")

	paddedAddress := utils.GetTxAddress(service.ToAddress)

	paddedAmount := utils.GetTxAmount(service.TokenID)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAmount...)
	data = append(data, paddedAddress...)

	fmt.Printf("service +%v", service)

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
			Code:  405,
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: signedTx.Hash(),
	}
}
