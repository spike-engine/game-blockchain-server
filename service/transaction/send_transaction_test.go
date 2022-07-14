package transaction

import (
	"fmt"
	"game-blockchain-server/service/signature"
	"game-blockchain-server/utils"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
	"math/big"
	"testing"
)

func TestSendTransaction(t *testing.T) {
	toAddress := common.HexToAddress("0xE88a42c47928818E5775fb3cd076792353bE938b")

	transferFnSignature := []byte("mint(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	amount := new(big.Int)
	//amount.SetString("1000000000000000000", 10) // 1000 tokens
	amount.SetString("100000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAmount...)
	data = append(data, paddedAddress...)

	spikeTx := &utils.SpikeTx{
		Data: data,
		To:   "0x3ebdAD182ea4F8D389c1EcDfeb3584F6bF416fd3",
	}
	transaction, err := spikeTx.ConstructionTransaction()
	if err != nil {
		fmt.Println("construction err :", err)
	}

	SignTxService := &signature.SignTxService{
		PassWord: "980125fyw",
		TX:       transaction,
	}

	res, err := SignTxService.SignSeparateTX()
	fmt.Printf("signedTX: +%v", res.Hash())

	Broad := &BroadcastService{
		SignedTX: res,
	}

	err = Broad.SendTransaction()

	fmt.Println(err)
}
