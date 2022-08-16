package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"privacy/x/privacy/types"
)

func (k Keeper) TxPrivacyDataAll(c context.Context, req *types.QueryAllTxPrivacyDataRequest) (*types.QueryAllTxPrivacyDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var txPrivacyDatas []types.TxPrivacyData
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	txPrivacyDataStore := prefix.NewStore(store, types.KeyPrefix(types.TxPrivacyDataKeyPrefix))

	pageRes, err := query.Paginate(txPrivacyDataStore, req.Pagination, func(key []byte, value []byte) error {
		var txPrivacyData types.TxPrivacyData
		if err := k.cdc.Unmarshal(value, &txPrivacyData); err != nil {
			return err
		}

		txPrivacyDatas = append(txPrivacyDatas, txPrivacyData)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTxPrivacyDataResponse{TxPrivacyData: txPrivacyDatas, Pagination: pageRes}, nil
}

func (k Keeper) TxPrivacyData(c context.Context, req *types.QueryGetTxPrivacyDataRequest) (*types.QueryGetTxPrivacyDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTxPrivacyData(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTxPrivacyDataResponse{TxPrivacyData: val}, nil
}
