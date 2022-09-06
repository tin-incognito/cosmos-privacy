package cli

import (
	"fmt"
	"strconv"
	"strings"

	"privacy/x/privacy/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer [private_key] [payment_infos]",
		Short: "Broadcast message transfer",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			privateKey := args[0]
			paymentInfos := []*types.MsgTransfer_PaymentInfo{}
			infos := strings.Split(args[1], ",")
			for _, v := range infos {
				temp := strings.Split(v, "-")
				paymentInfo := &types.MsgTransfer_PaymentInfo{}
				for i, value := range temp {
					if i == 0 {
						paymentInfo.PaymentAddress = value
					} else if i == 1 {
						amount, err := strconv.ParseUint(value, 10, 64)
						if err != nil {
							return err
						}
						paymentInfo.Amount = amount
					} else if i == 2 {
						paymentInfo.Info = []byte(value)
					} else {
						return fmt.Errorf("Invalid format payment infos %s", v)
					}
				}
				paymentInfos = append(paymentInfos, paymentInfo)
			}

			msg := types.NewMsgTransfer(
				clientCtx.GetFromAddress().String(),
				privateKey, paymentInfos,
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
