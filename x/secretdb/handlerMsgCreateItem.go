package secretdb

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/shunail2029/secretdb/x/secretdb/types"
	"github.com/shunail2029/secretdb/x/secretdb/keeper"
)

func handleMsgCreateItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateItem) (*sdk.Result, error) {
	var item = types.Item{
		Creator: msg.Creator,
		ID:      msg.ID,
	}
	k.CreateItem(ctx, item)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
