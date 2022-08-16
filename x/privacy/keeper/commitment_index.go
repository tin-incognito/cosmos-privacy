package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"privacy/x/privacy/types"
)

// SetCommitmentIndex set a specific commitmentIndex in the store from its index
func (k Keeper) SetCommitmentIndex(ctx sdk.Context, commitmentIndex types.CommitmentIndex) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CommitmentIndexKeyPrefix))
	b := k.cdc.MustMarshal(&commitmentIndex)
	store.Set(types.CommitmentIndexKey(
		commitmentIndex.Index,
	), b)
}

// GetCommitmentIndex returns a commitmentIndex from its index
func (k Keeper) GetCommitmentIndex(
	ctx sdk.Context,
	index string,

) (val types.CommitmentIndex, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CommitmentIndexKeyPrefix))

	b := store.Get(types.CommitmentIndexKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCommitmentIndex removes a commitmentIndex from the store
func (k Keeper) RemoveCommitmentIndex(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CommitmentIndexKeyPrefix))
	store.Delete(types.CommitmentIndexKey(
		index,
	))
}

// GetAllCommitmentIndex returns all commitmentIndex
func (k Keeper) GetAllCommitmentIndex(ctx sdk.Context) (list []types.CommitmentIndex) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CommitmentIndexKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CommitmentIndex
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
