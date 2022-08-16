package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"privacy/x/privacy/keeper"
	"privacy/x/privacy/types"
)

func SimulateMsgCreateTx(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateTx{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateTx simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateTx simulation not implemented"), nil, nil
	}
}
