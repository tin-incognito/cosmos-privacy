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

func (k Keeper) CommitmentIndexAll(c context.Context, req *types.QueryAllCommitmentIndexRequest) (*types.QueryAllCommitmentIndexResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var commitmentIndexs []types.CommitmentIndex
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	commitmentIndexStore := prefix.NewStore(store, types.KeyPrefix(types.CommitmentIndexKeyPrefix))

	pageRes, err := query.Paginate(commitmentIndexStore, req.Pagination, func(key []byte, value []byte) error {
		var commitmentIndex types.CommitmentIndex
		if err := k.cdc.Unmarshal(value, &commitmentIndex); err != nil {
			return err
		}

		commitmentIndexs = append(commitmentIndexs, commitmentIndex)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCommitmentIndexResponse{CommitmentIndex: commitmentIndexs, Pagination: pageRes}, nil
}

func (k Keeper) CommitmentIndex(c context.Context, req *types.QueryGetCommitmentIndexRequest) (*types.QueryGetCommitmentIndexResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCommitmentIndex(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCommitmentIndexResponse{CommitmentIndex: val}, nil
}
