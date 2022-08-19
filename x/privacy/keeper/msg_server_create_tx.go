package keeper

import (
	"context"

	"privacy/x/privacy/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateTx(goCtx context.Context, msg *types.MsgCreateTx) (*types.MsgCreateTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.setPrivacyData(ctx, msg.Value)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateTxResponse{}, nil
}
