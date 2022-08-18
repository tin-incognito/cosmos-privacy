package keeper

import (
	"context"

	"privacy/x/privacy/models"
	"privacy/x/privacy/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateTx(goCtx context.Context, msg *types.MsgCreateTx) (*types.MsgCreateTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// parse data
	serialNumbers, commitments, outputCoins, err := models.FetchDataFromTx(ctx, msg)
	if err != nil {
		return nil, err
	}

	// Store data
	// Store serialNumbers
	for _, serialNumber := range serialNumbers {
		k.SetSerialNumber(ctx, serialNumber)
	}

	for _, outputCoin := range outputCoins {
		for _, o := range outputCoin {
			k.SetOutputCoin(ctx, o)
		}
	}

	for _, commitment := range commitments {
		for _, c := range commitment {
			k.SetCommitment(ctx, c)
		}
	}

	// store one time address
	onetimeAddress := types.OnetimeAddress{}
	k.SetOnetimeAddress(ctx, onetimeAddress)

	return &types.MsgCreateTxResponse{}, nil
}
