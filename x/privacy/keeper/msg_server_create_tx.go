package keeper

import (
	"context"

	"privacy/x/privacy/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateTx(goCtx context.Context, msg *types.MsgCreateTx) (*types.MsgCreateTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_ = ctx

	return &types.MsgCreateTxResponse{}, nil
}
