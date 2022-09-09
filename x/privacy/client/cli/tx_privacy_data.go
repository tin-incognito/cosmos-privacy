package cli

import (
	"strconv"

	"privacy/x/privacy/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPrivacyData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "privacy-data",
		Short: "Broadcast message PrivacyData",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			/*msg := types.NewMsgPrivacyData(*/
			/*clientCtx.GetFromAddress().String(),*/
			/*)*/
			msg := &types.MsgPrivacyData{}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
