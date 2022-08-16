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

func (k Keeper) CommitmentAll(c context.Context, req *types.QueryAllCommitmentRequest) (*types.QueryAllCommitmentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var commitments []types.Commitment
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	commitmentStore := prefix.NewStore(store, types.KeyPrefix(types.CommitmentKeyPrefix))

	pageRes, err := query.Paginate(commitmentStore, req.Pagination, func(key []byte, value []byte) error {
		var commitment types.Commitment
		if err := k.cdc.Unmarshal(value, &commitment); err != nil {
			return err
		}

		commitments = append(commitments, commitment)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCommitmentResponse{Commitment: commitments, Pagination: pageRes}, nil
}

func (k Keeper) Commitment(c context.Context, req *types.QueryGetCommitmentRequest) (*types.QueryGetCommitmentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCommitment(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCommitmentResponse{Commitment: val}, nil
}