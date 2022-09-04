package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"privacy/x/privacy/types"
)

func CmdShowOutputCoinSerialNumber() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-output-coin-serial-number",
		Short: "shows OutputCoinSerialNumber",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetOutputCoinSerialNumberRequest{}

			res, err := queryClient.OutputCoinSerialNumber(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
