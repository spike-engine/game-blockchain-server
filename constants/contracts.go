package constants

import "golang.org/x/xerrors"

const (
	// SKK_CONTRACT_ADDRESS_TESTNET Governance coin ERC20
	SKK_CONTRACT_ADDRESS_TESTNET = ""
	SKK_CONTRACT_ADDRESS_MAINNET = ""

	// SKS_CONTRACT_ADDRESS_TESTNET Consumption coin ERC20
	SKS_CONTRACT_ADDRESS_TESTNET = ""
	SKS_CONTRACT_ADDRESS_MAINNET = ""

	// SOUL_CONTRACT_ADDRESS_TESTNET ERC721
	SOUL_CONTRACT_ADDRESS_TESTNET = ""
	SOUL_CONTRACT_ADDRESS_MAINNET = ""
)

func GetContractAddress(number int) (string, error) {
	switch number {
	case 1:
		return SKK_CONTRACT_ADDRESS_TESTNET, nil
	case 2:
		return SKK_CONTRACT_ADDRESS_MAINNET, nil
	case 3:
		return SKS_CONTRACT_ADDRESS_TESTNET, nil
	case 4:
		return SKS_CONTRACT_ADDRESS_MAINNET, nil
	case 5:
		return SOUL_CONTRACT_ADDRESS_TESTNET, nil
	case 6:
		return SOUL_CONTRACT_ADDRESS_MAINNET, nil

	}
	return "", xerrors.New("ContractNumber is not exist")
}
