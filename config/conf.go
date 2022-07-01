package config

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/joho/godotenv"
	"os"
)

func Init() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	ks := keystore.NewKeyStore(os.Getenv("KEY_DIR"), keystore.StandardScryptN, keystore.StandardScryptP)
	if len(ks.Accounts()) == 0 {
		_, err = ks.NewAccount("Creation password")
		if err != nil {
			panic(err)
		}
	}
}
