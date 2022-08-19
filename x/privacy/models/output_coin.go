package models

import (
	"math/big"
	"privacy/x/privacy/repos/coin"
)

func GenerateOutputCoin(amount big.Int, info []byte, otaReceiver coin.OTAReceiver) (*coin.Coin, error) {
	return coin.NewCoinFromAmountAndTxRandomBytes(amount, otaReceiver.PublicKey, &otaReceiver.TxRandom, info), nil
}
