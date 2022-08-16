package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"privacy/x/privacy/types"
)

func (k msgServer) CreateCommitmentIndex(goCtx context.Context, msg *types.MsgCreateCommitmentIndex) (*types.MsgCreateCommitmentIndexResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetCommitmentIndex(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var commitmentIndex = types.CommitmentIndex{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetCommitmentIndex(
		ctx,
		commitmentIndex,
	)
	return &types.MsgCreateCommitmentIndexResponse{}, nil
}

func (k msgServer) UpdateCommitmentIndex(goCtx context.Context, msg *types.MsgUpdateCommitmentIndex) (*types.MsgUpdateCommitmentIndexResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetCommitmentIndex(
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

	var commitmentIndex = types.CommitmentIndex{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetCommitmentIndex(ctx, commitmentIndex)

	return &types.MsgUpdateCommitmentIndexResponse{}, nil
}

func (k msgServer) DeleteCommitmentIndex(goCtx context.Context, msg *types.MsgDeleteCommitmentIndex) (*types.MsgDeleteCommitmentIndexResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetCommitmentIndex(
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

	k.RemoveCommitmentIndex(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteCommitmentIndexResponse{}, nil
}
