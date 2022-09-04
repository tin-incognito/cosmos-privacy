package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"privacy/x/privacy/types"
)

// SetOutputCoinSerialNumber set outputCoinSerialNumber in the store
func (k Keeper) SetOutputCoinSerialNumber(ctx sdk.Context, outputCoinSerialNumber types.OutputCoinSerialNumber) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OutputCoinSerialNumberKey))
	b := k.cdc.MustMarshal(&outputCoinSerialNumber)
	store.Set([]byte{0}, b)
}

// GetOutputCoinSerialNumber returns outputCoinSerialNumber
func (k Keeper) GetOutputCoinSerialNumber(ctx sdk.Context) (val types.OutputCoinSerialNumber, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OutputCoinSerialNumberKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOutputCoinSerialNumber removes outputCoinSerialNumber from the store
func (k Keeper) RemoveOutputCoinSerialNumber(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OutputCoinSerialNumberKey))
	store.Delete([]byte{0})
}
