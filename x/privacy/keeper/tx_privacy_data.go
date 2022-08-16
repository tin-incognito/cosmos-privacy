package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"privacy/x/privacy/types"
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
