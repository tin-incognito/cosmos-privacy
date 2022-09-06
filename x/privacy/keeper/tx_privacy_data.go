package keeper

import (
	"fmt"
	"math/big"
	"privacy/x/privacy/models"
	"privacy/x/privacy/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTxPrivacyData set a specific txPrivacyData in the store from its index
func (k Keeper) SetTxPrivacyData(ctx sdk.Context, txPrivacyData types.TxPrivacyData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxPrivacyDataKeyPrefix))
	b := k.cdc.MustMarshal(&txPrivacyData)
	store.Set(types.TxPrivacyDataKey(
		txPrivacyData.Index,
	), b)
}

// GetTxPrivacyData returns a txPrivacyData from its index
func (k Keeper) GetTxPrivacyData(
	ctx sdk.Context,
	index string,

) (val types.TxPrivacyData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxPrivacyDataKeyPrefix))

	b := store.Get(types.TxPrivacyDataKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTxPrivacyData removes a txPrivacyData from the store
func (k Keeper) RemoveTxPrivacyData(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxPrivacyDataKeyPrefix))
	store.Delete(types.TxPrivacyDataKey(
		index,
	))
}

// GetAllTxPrivacyData returns all txPrivacyData
func (k Keeper) GetAllTxPrivacyData(ctx sdk.Context) (list []types.TxPrivacyData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxPrivacyDataKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TxPrivacyData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) setPrivacyData(ctx sdk.Context, txPrivacyData []byte) error {
	outputCoinLength := big.NewInt(0)
	outputCoinSerialNumber, ok := k.GetOutputCoinSerialNumber(ctx)
	if ok {
		outputCoinLength.SetBytes(outputCoinSerialNumber.Value)
	}

	// parse data
	serialNumbers, commitments, outputCoins, otaCoins, newOutputCoinLength, err := models.FetchDataFromTx(ctx, txPrivacyData, *outputCoinLength)
	if err != nil {
		return err
	}

	// store serialNumber
	for _, serialNumber := range serialNumbers {
		if _, ok := k.GetSerialNumber(ctx, serialNumber.Index); ok {
			return fmt.Errorf("Duplicate serialNumber")
		}
		k.SetSerialNumber(ctx, serialNumber)
	}

	// store outputCoin
	for key, outputCoin := range outputCoins {
		otaCoin, found := otaCoins[key]
		if !found {
			return fmt.Errorf("Cannot find list otaCoin with key %v", key)
		}
		for i, o := range outputCoin {
			k.SetOutputCoin(ctx, o)
			if i >= len(otaCoin) {
				return fmt.Errorf("Cannot find otaCoin with key %v and index %v", key, i)
			}
			oa := otaCoin[i]
			k.SetOTACoin(ctx, oa)
		}
	}

	// store commitment
	for _, commitment := range commitments {
		for _, c := range commitment {
			if _, ok := k.GetCommitment(ctx, c.Index); ok {
				return fmt.Errorf("Duplicate commitment")
			}
			k.SetCommitment(ctx, c)
		}
	}

	// store output coin length
	k.SetOutputCoinSerialNumber(ctx, types.OutputCoinSerialNumber{Value: newOutputCoinLength.Bytes()})

	return nil
}
