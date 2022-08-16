package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"privacy/x/privacy/types"
)

func (k msgServer) CreateOutputCoin(goCtx context.Context, msg *types.MsgCreateOutputCoin) (*types.MsgCreateOutputCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetOutputCoin(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var outputCoin = types.OutputCoin{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetOutputCoin(
		ctx,
		outputCoin,
	)
	return &types.MsgCreateOutputCoinResponse{}, nil
}

func (k msgServer) UpdateOutputCoin(goCtx context.Context, msg *types.MsgUpdateOutputCoin) (*types.MsgUpdateOutputCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetOutputCoin(
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

	var outputCoin = types.OutputCoin{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetOutputCoin(ctx, outputCoin)

	return &types.MsgUpdateOutputCoinResponse{}, nil
}

func (k msgServer) DeleteOutputCoin(goCtx context.Context, msg *types.MsgDeleteOutputCoin) (*types.MsgDeleteOutputCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetOutputCoin(
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

	k.RemoveOutputCoin(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteOutputCoinResponse{}, nil
}
