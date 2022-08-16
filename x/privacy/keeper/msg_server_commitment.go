package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"privacy/x/privacy/types"
)

func (k msgServer) CreateCommitment(goCtx context.Context, msg *types.MsgCreateCommitment) (*types.MsgCreateCommitmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetCommitment(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var commitment = types.Commitment{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetCommitment(
		ctx,
		commitment,
	)
	return &types.MsgCreateCommitmentResponse{}, nil
}

func (k msgServer) UpdateCommitment(goCtx context.Context, msg *types.MsgUpdateCommitment) (*types.MsgUpdateCommitmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetCommitment(
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

	var commitment = types.Commitment{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetCommitment(ctx, commitment)

	return &types.MsgUpdateCommitmentResponse{}, nil
}

func (k msgServer) DeleteCommitment(goCtx context.Context, msg *types.MsgDeleteCommitment) (*types.MsgDeleteCommitmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetCommitment(
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

	k.RemoveCommitment(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteCommitmentResponse{}, nil
}
