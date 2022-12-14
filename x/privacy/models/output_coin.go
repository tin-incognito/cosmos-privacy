package models

import (
	"fmt"
	"privacy/x/privacy/repos/coin"
	"privacy/x/privacy/repos/key"
	"privacy/x/privacy/types"
	"sort"

	"github.com/incognitochain/go-incognito-sdk-v2/wallet"
)

type OutputCoin struct {
	index string
	value *coin.Coin
}

func GenerateOutputCoin(amount uint64, info []byte, otaReceiver coin.OTAReceiver) (*coin.Coin, error) {
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
	paymentInfos []*types.MsgTransfer_PaymentInfo, feePerKb uint64,
	metadata []byte,
) ([]*OutputCoin, []*key.PaymentInfo, uint64, error) {
	var res, remainCoins []*OutputCoin
	var resPaymentInfos []*key.PaymentInfo
	res, err := getCoinsByKeySet(coins, keySet)
	if err != nil {
		return nil, nil, 0, err
	}
	var candidateOutputCoinAmount uint64
	res, remainCoins, candidateOutputCoinAmount, err = chooseBestOutCoinsToSpent(res, amount)
	if err != nil {
		return nil, nil, 0, err
	}

	for _, info := range paymentInfos {
		paymentAddress := key.PaymentAddress{}
		keyWallet, err := wallet.Base58CheckDeserialize(info.PaymentAddress)
		if err != nil {
			return nil, nil, 0, err
		}
		err = paymentAddress.SetBytes(keyWallet.KeySet.PaymentAddress.Bytes())
		if err != nil {
			return nil, nil, 0, err
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
	fee := EstimateFee(feePerKb, len(coins), len(paymentInfos), metadata)
	needToPayFee := int64((amount + fee) - candidateOutputCoinAmount)
	// if not enough to pay fee
	if needToPayFee > 0 {
		if len(remainCoins) > 0 {
			candidateOutputCoinsForFee, _, _, err := chooseBestOutCoinsToSpent(remainCoins, uint64(needToPayFee))
			if err != nil {
				return nil, nil, 0, err
			}
			res = append(res, candidateOutputCoinsForFee...)
		}
	}

	if overBalanceAmount > 0 {
		lastPaymentInfo := resPaymentInfos[len(resPaymentInfos)-1]
		if lastPaymentInfo.PaymentAddress.String() == keySet.PaymentAddress.String() {
			temp := lastPaymentInfo.Amount - fee
			if temp > lastPaymentInfo.Amount {
				return nil, nil, 0, fmt.Errorf("out of range uint64")
			}
			lastPaymentInfo.Amount = temp
		}
		resPaymentInfos[len(resPaymentInfos)-1] = lastPaymentInfo
	}

	return res, resPaymentInfos, fee, nil

}

func getCoinsByKeySet(coins []types.OutputCoin, keySet key.KeySet) ([]*OutputCoin, error) {
	var res []*OutputCoin
	for _, outputCoin := range coins {
		o := &coin.Coin{}
		err := o.SetBytes(outputCoin.Value)
		if err != nil {
			return nil, err
		}
		if ok, _ := o.DoesCoinBelongToKeySet(&keySet); !ok {
			continue
		}
		oc := &OutputCoin{
			index: outputCoin.Index,
			value: o,
		}
		o, err = o.Decrypt(&keySet)
		if err != nil {
			return nil, err
		}
		res = append(res, oc)
	}
	return res, nil
}

func chooseBestOutCoinsToSpent(coins []*OutputCoin, amount uint64) ([]*OutputCoin, []*OutputCoin, uint64, error) {
	resultOutputCoins := []*OutputCoin{}
	remainOutputCoins := []*OutputCoin{}
	totalResultOutputCoinAmount := uint64(0)

	// either take the smallest coins, or a single largest one
	var outCoinOverLimit *OutputCoin
	outCoinsUnderLimit := make([]*OutputCoin, 0)
	for _, c := range coins {
		if c.value.GetValue() < amount {
			outCoinsUnderLimit = append(outCoinsUnderLimit, c)
		} else if outCoinOverLimit == nil {
			outCoinOverLimit = new(OutputCoin)
			*outCoinOverLimit = *c
		} else if outCoinOverLimit.value.GetValue() > c.value.GetValue() {
			remainOutputCoins = append(remainOutputCoins, c)
		} else {
			remainOutputCoins = append(remainOutputCoins, outCoinOverLimit)
			*outCoinOverLimit = *c
		}
	}
	sort.Slice(outCoinsUnderLimit, func(i, j int) bool {
		return outCoinsUnderLimit[i].value.GetValue() < outCoinsUnderLimit[j].value.GetValue()
	})
	for _, outCoin := range outCoinsUnderLimit {
		if totalResultOutputCoinAmount < amount {
			totalResultOutputCoinAmount += outCoin.value.GetValue()
			resultOutputCoins = append(resultOutputCoins, outCoin)
		} else {
			remainOutputCoins = append(remainOutputCoins, outCoin)
		}
	}
	if outCoinOverLimit != nil && (outCoinOverLimit.value.GetValue() > 2*amount || totalResultOutputCoinAmount < amount) {
		remainOutputCoins = append(remainOutputCoins, resultOutputCoins...)
		resultOutputCoins = []*OutputCoin{outCoinOverLimit}
		totalResultOutputCoinAmount = outCoinOverLimit.value.GetValue()
	} else if outCoinOverLimit != nil {
		remainOutputCoins = append(remainOutputCoins, outCoinOverLimit)
	}
	if totalResultOutputCoinAmount < amount {
		return resultOutputCoins, remainOutputCoins, totalResultOutputCoinAmount, fmt.Errorf("Not enought coin")
	} else {
		return resultOutputCoins, remainOutputCoins, totalResultOutputCoinAmount, nil
	}
}
