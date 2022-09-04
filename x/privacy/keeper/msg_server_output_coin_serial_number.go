package keeper

import (
	"context"

	"privacy/x/privacy/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateOutputCoinSerialNumber(goCtx context.Context, msg *types.MsgCreateOutputCoinSerialNumber) (*types.MsgCreateOutputCoinSerialNumberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetOutputCoinSerialNumber(ctx)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var outputCoinSerialNumber = types.OutputCoinSerialNumber{}

	k.SetOutputCoinSerialNumber(
		ctx,
		outputCoinSerialNumber,
	)
	return &types.MsgCreateOutputCoinSerialNumberResponse{}, nil
}

func (k msgServer) UpdateOutputCoinSerialNumber(goCtx context.Context, msg *types.MsgUpdateOutputCoinSerialNumber) (*types.MsgUpdateOutputCoinSerialNumberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	_, isFound := k.GetOutputCoinSerialNumber(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	var outputCoinSerialNumber = types.OutputCoinSerialNumber{}

	k.SetOutputCoinSerialNumber(ctx, outputCoinSerialNumber)

	return &types.MsgUpdateOutputCoinSerialNumberResponse{}, nil
}

func (k msgServer) DeleteOutputCoinSerialNumber(goCtx context.Context, msg *types.MsgDeleteOutputCoinSerialNumber) (*types.MsgDeleteOutputCoinSerialNumberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	_, isFound := k.GetOutputCoinSerialNumber(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	k.RemoveOutputCoinSerialNumber(ctx)

	return &types.MsgDeleteOutputCoinSerialNumberResponse{}, nil
}
