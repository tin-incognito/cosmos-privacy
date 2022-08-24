package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/incognitochain/go-incognito-sdk-v2/incclient"
	"github.com/incognitochain/go-incognito-sdk-v2/wallet"
)

type accountInfo struct {
	Index int
	*incclient.KeyInfo
}

func ImportAccount(mnemonic string, numAccounts int) error {
	w, err := wallet.NewMasterKeyFromMnemonic(mnemonic)
	if err != nil {
		return err
	}

	accounts := make([]*accountInfo, 0)
	genCount := 0
	index := 1
	for {
		if genCount == numAccounts {
			break
		}
		childKey, err := w.DeriveChild(uint32(index))
		if err != nil {
			return err
		}
		privateKey := childKey.Base58CheckSerialize(wallet.PrivateKeyType)
		info, err := incclient.GetAccountInfoFromPrivateKey(privateKey)
		if err != nil {
			return err
		}
		accounts = append(accounts, &accountInfo{Index: index, KeyInfo: info})
		genCount++
		index++
	}
	data, err := json.Marshal(accounts)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func main() {
	mnemonic := os.Args[1]
	mnemonic = strings.Replace(mnemonic, "-", " ", -1)
	numAccountsArgs := os.Args[2]
	numAccounts, err := strconv.Atoi(numAccountsArgs)
	if err != nil {
		panic(err)
	}

	err = ImportAccount(mnemonic, numAccounts)
	if err != nil {
		panic(err)
	}
}
