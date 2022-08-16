package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"privacy/x/privacy/types"
)

func (k msgServer) CreateOnetimeAddress(goCtx context.Context, msg *types.MsgCreateOnetimeAddress) (*types.MsgCreateOnetimeAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetOnetimeAddress(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var onetimeAddress = types.OnetimeAddress{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetOnetimeAddress(
		ctx,
		onetimeAddress,
	)
	return &types.MsgCreateOnetimeAddressResponse{}, nil
}

func (k msgServer) UpdateOnetimeAddress(goCtx context.Context, msg *types.MsgUpdateOnetimeAddress) (*types.MsgUpdateOnetimeAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetOnetimeAddress(
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

	var onetimeAddress = types.OnetimeAddress{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetOnetimeAddress(ctx, onetimeAddress)

	return &types.MsgUpdateOnetimeAddressResponse{}, nil
}

func (k msgServer) DeleteOnetimeAddress(goCtx context.Context, msg *types.MsgDeleteOnetimeAddress) (*types.MsgDeleteOnetimeAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetOnetimeAddress(
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

	k.RemoveOnetimeAddress(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteOnetimeAddressResponse{}, nil
}
