package keeper

import (
	"context"

	"privacy/x/privacy/types"
)

func (k msgServer) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {

	return &types.MsgTransferResponse{
		Msg: "Success",
	}, nil
}
