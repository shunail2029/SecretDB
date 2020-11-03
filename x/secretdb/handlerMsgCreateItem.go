package secretdb

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/shunail2029/secretdb/x/secretdb/keeper"
	"github.com/shunail2029/secretdb/x/secretdb/types"
)

// Handle a message to create item
// TODO: use return value of CreateItem
func handleMsgCreateItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateItem) (*sdk.Result, error) {
	var item = types.Item{
		Owner: msg.Owner,
		Data:  msg.Data,
	}
	k.CreateItem(ctx, item)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
