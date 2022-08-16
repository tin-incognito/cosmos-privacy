package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"privacy/x/privacy/types"
)

func CmdListTxPrivacyData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-tx-privacy-data",
		Short: "list all TxPrivacyData",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTxPrivacyDataRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TxPrivacyDataAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowTxPrivacyData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-tx-privacy-data [index]",
		Short: "shows a TxPrivacyData",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetTxPrivacyDataRequest{
				Index: argIndex,
			}

			res, err := queryClient.TxPrivacyData(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
