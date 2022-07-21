package utils

import (
	"bytes"
	"encoding/binary"
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
	stringLength := len([]byte(s))
	var result []byte
	result = append(result, common.LeftPadBytes(IntToBytes(stringLength), 32)...)
	result = append(result, common.LeftPadBytes([]byte(s), 32)...)
	return result
}

func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func GetOffset(offset int) []byte {
	return common.LeftPadBytes(IntToBytes(offset*32), 32)
}
