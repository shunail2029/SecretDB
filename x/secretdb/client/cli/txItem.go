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
	"github.com/shunail2029/secretdb/x/secretdb/types"
)

// GetCmdCreateItem ...
func GetCmdCreateItem(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-item [data]",
		Short: "Creates a new item",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			var data bson.M
			err := bson.Unmarshal([]byte(args[0]), &data)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateItem(cliCtx.GetFromAddress(), data)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdSetItem ...
func GetCmdSetItem(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-item [filter] [update] ",
		Short: "Set a new item",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var filter bson.D
			err := bson.Unmarshal([]byte(args[0]), &filter)
			if err != nil {
				return err
			}
			var update bson.D
			err = bson.Unmarshal([]byte(args[1]), &update)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetItem(cliCtx.GetFromAddress(), filter, update)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdSetItems ...
func GetCmdSetItems(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-items [filter] [update] ",
		Short: "Set some new items",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var filter bson.D
			err := bson.Unmarshal([]byte(args[0]), &filter)
			if err != nil {
				return err
			}
			var update bson.D
			err = bson.Unmarshal([]byte(args[1]), &update)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetItems(cliCtx.GetFromAddress(), filter, update)
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
			var filter bson.D
			err := bson.Unmarshal([]byte(args[0]), &filter)
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
			var filter bson.D
			err := bson.Unmarshal([]byte(args[0]), &filter)
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
