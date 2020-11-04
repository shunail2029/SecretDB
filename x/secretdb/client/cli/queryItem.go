package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/shunail2029/SecretDB/x/secretdb/types"
	"github.com/spf13/cobra"
)

// GetCmdGetItem ...
func GetCmdGetItem(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-item [filter]",
		Short: "Query a item by filter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			filter := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetItem, filter), nil)
			if err != nil {
				fmt.Printf("could not resolve item %s \n%s\n", filter, err.Error())

				return nil
			}

			var out types.Item
			cdc.MustUnmarshalJSON(res, &out) // TODO: check MustUnmarshalJSON
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdGetItems ...
func GetCmdGetItems(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-items [filter]",
		Short: "Query some items by filter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			filter := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetItem, filter), nil)
			if err != nil {
				fmt.Printf("could not resolve item %s \n%s\n", filter, err.Error())

				return nil
			}

			var out types.Item
			cdc.MustUnmarshalJSON(res, &out) // TODO: check MustUnmarshalJSON
			return cliCtx.PrintOutput(out)
		},
	}
}
