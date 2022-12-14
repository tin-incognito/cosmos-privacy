package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"privacy/x/privacy/types"
)

var _ = strconv.Itoa(0)

func CmdBalance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balance [private-key]",
		Short: "Query Balance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPrivateKey := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryBalanceRequest{

				PrivateKey: reqPrivateKey,
			}

			res, err := queryClient.Balance(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
