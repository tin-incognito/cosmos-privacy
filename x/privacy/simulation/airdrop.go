package simulation

import (
	"math/rand"

	"privacy/x/privacy/keeper"
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
		//simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAirdrop{
			//Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Airdrop simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Airdrop simulation not implemented"), nil, nil
	}
}
