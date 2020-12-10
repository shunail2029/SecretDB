package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/shunail2029/SecretDB/x/secretdb/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group secretdb queries under a subcommand
	secretdbQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmds := flags.GetCommands(
		GetCmdGetItem(queryRoute, cdc),
		GetCmdGetItems(queryRoute, cdc),
	)
	for _, c := range cmds {
		c.Flags().String(flags.FlagFrom, "", "Name or address of private key with which to sign")
	}

	secretdbQueryCmd.AddCommand(cmds...)

	return secretdbQueryCmd
}
