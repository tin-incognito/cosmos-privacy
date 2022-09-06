package models

import (
	"fmt"
	"math/big"
	"privacy/x/privacy/repos/coin"
	"privacy/x/privacy/repos/key"
	"privacy/x/privacy/types"
	"sort"

	"github.com/incognitochain/go-incognito-sdk-v2/wallet"
)

func GenerateOutputCoin(amount big.Int, info []byte, otaReceiver coin.OTAReceiver) (*coin.Coin, error) {
	return coin.NewCoinFromAmountAndTxRandomBytes(amount, otaReceiver.PublicKey, &otaReceiver.TxRandom, info), nil
}

func GenerateOutputCoinsByPaymentInfos(paymentInfos []*key.PaymentInfo) ([]*coin.Coin, error) {
	res := make([]*coin.Coin, len(paymentInfos))
	for index, info := range paymentInfos {
		c, err := coin.NewCoinFromPaymentInfo(info)
		if err != nil {
			return nil, err
		}
		res[index] = c
	}

	return res, nil
}

func chooseCoinsByKeySet(
	coins []types.OutputCoin, keySet key.KeySet, amount uint64,
	paymentInfos []*types.MsgTransfer_PaymentInfo,
) ([]*coin.Coin, []*key.PaymentInfo, error) {
	var res []*coin.Coin
	var resPaymentInfos []*key.PaymentInfo
	res, err := getCoinsByKeySet(coins, keySet)
	if err != nil {
		return nil, nil, err
	}
	var candidateOutputCoinAmount uint64
	res, candidateOutputCoinAmount, err = chooseBestOutCoinsToSpent(res, amount)
	if err != nil {
		return nil, nil, err
	}

	for _, info := range paymentInfos {
		paymentAddress := key.PaymentAddress{}
		keyWallet, err := wallet.Base58CheckDeserialize(info.PaymentAddress)
		if err != nil {
			return nil, nil, err
		}
		err = paymentAddress.SetBytes(keyWallet.KeySet.PaymentAddress.Bytes())
		if err != nil {
			return nil, nil, err
		}
		resPaymentInfos = append(resPaymentInfos, &key.PaymentInfo{
			Amount:         info.Amount,
			Message:        info.Info,
			PaymentAddress: paymentAddress,
		})
	}

	overBalanceAmount := candidateOutputCoinAmount - amount
	if overBalanceAmount > 0 {
		// add more into output for estimate fee
		resPaymentInfos = append(resPaymentInfos, &key.PaymentInfo{
			PaymentAddress: keySet.PaymentAddress,
			Amount:         overBalanceAmount,
		})
	}
	return res, resPaymentInfos, nil

}

func getCoinsByKeySet(coins []types.OutputCoin, keySet key.KeySet) ([]*coin.Coin, error) {
	var res []*coin.Coin
	for _, outputCoin := range coins {
		o := &coin.Coin{}
		err := o.SetBytes(outputCoin.Value)
		if err != nil {
			return nil, err
		}
		if ok, _ := o.DoesCoinBelongToKeySet(&keySet); !ok {
			continue
		}
		o, err = o.Decrypt(&keySet)
		if err != nil {
			return nil, err
		}
		res = append(res, o)
	}
	return res, nil
}

func chooseBestOutCoinsToSpent(coins []*coin.Coin, amount uint64) ([]*coin.Coin, uint64, error) {
	resultOutputCoins := []*coin.Coin{}
	totalResultOutputCoinAmount := uint64(0)

	// either take the smallest coins, or a single largest one
	var outCoinOverLimit *coin.Coin
	outCoinsUnderLimit := make([]*coin.Coin, 0)
	for _, c := range coins {
		if c.GetValue() < amount {
			outCoinsUnderLimit = append(outCoinsUnderLimit, c)
		} else if outCoinOverLimit == nil {
			outCoinOverLimit = new(coin.Coin)
			*outCoinOverLimit = *c
		} else if outCoinOverLimit.GetValue() <= c.GetValue() {
			outCoinOverLimit = new(coin.Coin)
			*outCoinOverLimit = *c
		}
	}
	sort.Slice(outCoinsUnderLimit, func(i, j int) bool {
		return outCoinsUnderLimit[i].GetValue() < outCoinsUnderLimit[j].GetValue()
	})
	for _, outCoin := range outCoinsUnderLimit {
		if totalResultOutputCoinAmount < amount {
			totalResultOutputCoinAmount += outCoin.GetValue()
			resultOutputCoins = append(resultOutputCoins, outCoin)
		}
	}
	if outCoinOverLimit != nil && (outCoinOverLimit.GetValue() > 2*amount || totalResultOutputCoinAmount < amount) {
		resultOutputCoins = []*coin.Coin{outCoinOverLimit}
		totalResultOutputCoinAmount = outCoinOverLimit.GetValue()
	}
	if totalResultOutputCoinAmount < amount {
		return resultOutputCoins, totalResultOutputCoinAmount, fmt.Errorf("Not enought coin")
	} else {
		return resultOutputCoins, totalResultOutputCoinAmount, nil
	}
}
