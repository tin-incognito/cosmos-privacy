package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"privacy/x/privacy/types"
)

func (k Keeper) OutputCoinSerialNumber(c context.Context, req *types.QueryGetOutputCoinSerialNumberRequest) (*types.QueryGetOutputCoinSerialNumberResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetOutputCoinSerialNumber(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetOutputCoinSerialNumberResponse{OutputCoinSerialNumber: val}, nil
}
