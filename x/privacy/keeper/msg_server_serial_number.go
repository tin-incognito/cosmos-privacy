package keeper

import (
	"context"

	"privacy/x/privacy/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateSerialNumber(goCtx context.Context, msg *types.MsgCreateSerialNumber) (*types.MsgCreateSerialNumberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetSerialNumber(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var serialNumber = types.SerialNumber{
		Index: msg.Index,
	}

	k.SetSerialNumber(
		ctx,
		serialNumber,
	)
	return &types.MsgCreateSerialNumberResponse{}, nil
}

func (k msgServer) UpdateSerialNumber(goCtx context.Context, msg *types.MsgUpdateSerialNumber) (*types.MsgUpdateSerialNumberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	_, isFound := k.GetSerialNumber(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	var serialNumber = types.SerialNumber{
		Index: msg.Index,
	}

	k.SetSerialNumber(ctx, serialNumber)

	return &types.MsgUpdateSerialNumberResponse{}, nil
}

func (k msgServer) DeleteSerialNumber(goCtx context.Context, msg *types.MsgDeleteSerialNumber) (*types.MsgDeleteSerialNumberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	_, isFound := k.GetSerialNumber(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	k.RemoveSerialNumber(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteSerialNumberResponse{}, nil
}
