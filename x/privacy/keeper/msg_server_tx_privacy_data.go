package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"privacy/x/privacy/types"
)

func (k msgServer) CreateTxPrivacyData(goCtx context.Context, msg *types.MsgCreateTxPrivacyData) (*types.MsgCreateTxPrivacyDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetTxPrivacyData(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var txPrivacyData = types.TxPrivacyData{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetTxPrivacyData(
		ctx,
		txPrivacyData,
	)
	return &types.MsgCreateTxPrivacyDataResponse{}, nil
}

func (k msgServer) UpdateTxPrivacyData(goCtx context.Context, msg *types.MsgUpdateTxPrivacyData) (*types.MsgUpdateTxPrivacyDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTxPrivacyData(
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

	var txPrivacyData = types.TxPrivacyData{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetTxPrivacyData(ctx, txPrivacyData)

	return &types.MsgUpdateTxPrivacyDataResponse{}, nil
}

func (k msgServer) DeleteTxPrivacyData(goCtx context.Context, msg *types.MsgDeleteTxPrivacyData) (*types.MsgDeleteTxPrivacyDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTxPrivacyData(
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

	k.RemoveTxPrivacyData(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteTxPrivacyDataResponse{}, nil
}
