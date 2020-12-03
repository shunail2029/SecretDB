package cli

import (
	"bufio"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/shunail2029/SecretDB/x/secretdb/types"
)

// GetCmdStoreItem ...
func GetCmdStoreItem(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "store-item [data]",
		Short: "Stores a new item",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var data bson.M
			err := bson.UnmarshalExtJSON([]byte(args[0]), false, &data)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgStoreItem(cliCtx.GetFromAddress(), data)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdUpdateItem ...
func GetCmdUpdateItem(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "update-item [filter] [update]",
		Short: "Update a new item",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var filter bson.M
			err := bson.UnmarshalExtJSON([]byte(args[0]), true, &filter)
			if err != nil {
				return err
			}
			var update bson.M
			err = bson.UnmarshalExtJSON([]byte(args[1]), true, &update)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgUpdateItem(cliCtx.GetFromAddress(), filter, update)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdUpdateItems ...
func GetCmdUpdateItems(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "update-items [filter] [update]",
		Short: "Update some new items",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var filter bson.M
			err := bson.UnmarshalExtJSON([]byte(args[0]), true, &filter)
			if err != nil {
				return err
			}
			var update bson.M
			err = bson.UnmarshalExtJSON([]byte(args[1]), true, &update)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgUpdateItems(cliCtx.GetFromAddress(), filter, update)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdDeleteItem ...
func GetCmdDeleteItem(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-item [filter]",
		Short: "Delete a item by Filter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var filter bson.M
			err := bson.UnmarshalExtJSON([]byte(args[0]), true, &filter)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteItem(cliCtx.GetFromAddress(), filter)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdDeleteItems ...
func GetCmdDeleteItems(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-items [filter]",
		Short: "Delete some items by Filter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var filter bson.M
			err := bson.UnmarshalExtJSON([]byte(args[0]), true, &filter)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteItems(cliCtx.GetFromAddress(), filter)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
