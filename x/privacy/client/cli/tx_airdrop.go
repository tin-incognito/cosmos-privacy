package cli

import (
	"encoding/json"
	"strconv"

	"privacy/x/privacy/common"
	"privacy/x/privacy/models"
	"privacy/x/privacy/repos/coin"

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
			amount, err := strconv.ParseUint(argAmount, 10, 64)
			if err != nil {
				return err
			}
			otaReceiver := coin.OTAReceiver{}
			err = otaReceiver.FromString(argOTAReceiver)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			type Message struct {
				OtaReceiver string `json:"ota_receiver"`
				Amount      uint64 `json:"amount"`
				Info        []byte `json:"info"`
			}
			m := Message{
				OtaReceiver: argOTAReceiver,
				Amount:      amount,
				Info:        nil,
			}

			msgBytes, err := json.Marshal(m)
			if err != nil {
				return err
			}
			hash := common.HashH(msgBytes)

			msg, err := models.BuildMintTx(otaReceiver, amount, nil, nil, hash)
			if err != nil {
				return err
			}
			msg.Creator = clientCtx.GetFromAddress().String()

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
