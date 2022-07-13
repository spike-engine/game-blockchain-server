package mint

import (
	"game-blockchain-server/constants"
	"game-blockchain-server/serializer"
	"game-blockchain-server/service/signature"
	tx "game-blockchain-server/service/transaction"
	"game-blockchain-server/utils"
	"os"
)

type MintERC721Service struct {
	TokenID string `json:"token_id,omitempty"`
	//  toAddress is nft owner, Can be the owner or administrator of the contract
	ToAddress string `json:"to_address,omitempty"`
}

func (service *MintERC721Service) MintSoul() serializer.Response {
	methodID := utils.GetTxMethodName("mint(address,uint256)")

	paddedAddress := utils.GetTxAddress(service.ToAddress)

	paddedAmount := utils.GetTxAmount(service.TokenID)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAmount...)
	data = append(data, paddedAddress...)

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
