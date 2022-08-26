package models

import (
	"math/big"
	"privacy/x/privacy/repos/coin"
)

func GenerateOutputCoin(amount big.Int, info []byte, otaReceiver coin.OTAReceiver) (*coin.Coin, error) {
	return coin.NewCoinFromAmountAndTxRandomBytes(amount, otaReceiver.PublicKey, &otaReceiver.TxRandom, info), nil
}

func GetBalance() (uint64, error) {
	var res uint64
	return res, nil
}
