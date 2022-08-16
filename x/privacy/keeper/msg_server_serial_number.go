package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"privacy/x/privacy/types"
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
		Creator: msg.Creator,
		Index:   msg.Index,
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
	valFound, isFound := k.GetSerialNumber(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var serialNumber = types.SerialNumber{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetSerialNumber(ctx, serialNumber)

	return &types.MsgUpdateSerialNumberResponse{}, nil
}

func (k msgServer) DeleteSerialNumber(goCtx context.Context, msg *types.MsgDeleteSerialNumber) (*types.MsgDeleteSerialNumberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetSerialNumber(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveSerialNumber(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteSerialNumberResponse{}, nil
}
