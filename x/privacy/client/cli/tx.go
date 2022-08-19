package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"privacy/x/privacy/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateSerialNumber())
	cmd.AddCommand(CmdUpdateSerialNumber())
	cmd.AddCommand(CmdDeleteSerialNumber())
	cmd.AddCommand(CmdCreateOutputCoin())
	cmd.AddCommand(CmdUpdateOutputCoin())
	cmd.AddCommand(CmdDeleteOutputCoin())
	cmd.AddCommand(CmdCreateCommitment())
	cmd.AddCommand(CmdUpdateCommitment())
	cmd.AddCommand(CmdDeleteCommitment())
	cmd.AddCommand(CmdCreateCommitmentIndex())
	cmd.AddCommand(CmdUpdateCommitmentIndex())
	cmd.AddCommand(CmdDeleteCommitmentIndex())
	cmd.AddCommand(CmdCreateToken())
	cmd.AddCommand(CmdUpdateToken())
	cmd.AddCommand(CmdDeleteToken())
	cmd.AddCommand(CmdCreateOnetimeAddress())
	cmd.AddCommand(CmdUpdateOnetimeAddress())
	cmd.AddCommand(CmdDeleteOnetimeAddress())
	cmd.AddCommand(CmdCreateTx())
	cmd.AddCommand(CmdCreateTxPrivacyData())
	cmd.AddCommand(CmdUpdateTxPrivacyData())
	cmd.AddCommand(CmdDeleteTxPrivacyData())
	cmd.AddCommand(CmdAirdrop())
	// this line is used by starport scaffolding # 1

	return cmd
}
