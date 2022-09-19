package simulation

import (
	"math/rand"

	"privacy/x/privacy/common"
	"privacy/x/privacy/keeper"
	"privacy/x/privacy/models"
	"privacy/x/privacy/repos/coin"
	"privacy/x/privacy/repos/key"
	"privacy/x/privacy/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAirdrop(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		ks := &key.KeySet{}
		ks = ks.GenerateKey([]byte("account 0"))
		otaReceiver := coin.OTAReceiver{}
		err := otaReceiver.FromAddress(ks.PaymentAddress)
		amount := r.Uint64()
		h := common.HashH(common.Uint32ToBytes(r.Uint32()))
		msg, err := models.BuildMintTx(otaReceiver, amount, nil, nil, h)
		if err != nil {
			panic(err)
		}

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Airdrop simulation not implemented"), nil, nil
	}
}
