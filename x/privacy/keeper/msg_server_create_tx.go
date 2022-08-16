package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"privacy/x/privacy/types"
)

func (k msgServer) CreateTx(goCtx context.Context, msg *types.MsgCreateTx) (*types.MsgCreateTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateTxResponse{}, nil
}
