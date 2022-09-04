package cli

import (
	"fmt"
	"math/big"
	"strconv"

	"privacy/x/privacy/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAirdrop() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "airdrop [ota_receiver] [amount] [info]",
		Short: "Broadcast message airdrop",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOTAReceiver := args[0]
			argAmount := args[1]
			amount, ok := big.NewInt(0).SetString(argAmount, 10)
			if !ok {
				err = fmt.Errorf("Invalid amount")
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAirdrop(
				argOTAReceiver,
				amount.Bytes(),
				nil,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
