package keeper

import (
	"context"

	"privacy/x/privacy/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Airdrop(goCtx context.Context, msg *types.MsgAirdrop) (*types.MsgAirdropResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	return &types.MsgAirdropResponse{}, nil
}
