package models

import (
	"fmt"
	"privacy/x/privacy/repos/coin"
	"privacy/x/privacy/repos/key"
	"privacy/x/privacy/types"
)

func BalanceByKeySet(coins []types.OutputCoin, keySet key.KeySet) (uint64, error) {
	var res uint64
	for _, outputCoin := range coins {
		o := &coin.Coin{}
		err := o.SetBytes(outputCoin.Value)
		if err != nil {
			return 0, err
		}
		o, err = o.Decrypt(&keySet)
		if err == nil {
			temp := res + o.GetValue()
			if temp < res {
				return 0, fmt.Errorf("balance is out of range")
			}
			res = temp
		}
	}
	return res, nil
}
