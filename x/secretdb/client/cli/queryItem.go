package cli

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

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

			err := bson.UnmarshalExtJSON([]byte(args[0]), true, bson.M{})
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetItem, args[0]), nil)
			if err != nil {
				fmt.Printf("could not resolve item %s \n%s\n", args[0], err.Error())
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

			err := bson.UnmarshalExtJSON([]byte(args[0]), true, bson.M{})
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetItems, args[0]), nil)
			if err != nil {
				fmt.Printf("could not resolve item %s \n%s\n", args[0], err.Error())
				return nil
			}

			var out types.Item
			cdc.MustUnmarshalJSON(res, &out) // TODO: check MustUnmarshalJSON
			return cliCtx.PrintOutput(out)
		},
	}
}
