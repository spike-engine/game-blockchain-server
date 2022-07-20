package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
	"math/big"
)

func GetTxMethodName(method string) []byte {
	methodName := []byte(method)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(methodName)
	return hash.Sum(nil)[:4]
}

func GetTxAddress(address string) []byte {
	return common.LeftPadBytes(common.HexToAddress(address).Bytes(), 32)
}

func GetTxUint256(a string) []byte {
	amount := new(big.Int)
	amount.SetString(a, 10) // 1000 tokens
	return common.LeftPadBytes(amount.Bytes(), 32)
}

func GetTxString(s string) []byte {
	return common.LeftPadBytes([]byte(s), 32)
}
