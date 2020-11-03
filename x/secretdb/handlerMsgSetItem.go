package secretdb

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/shunail2029/secretdb/x/secretdb/keeper"
	"github.com/shunail2029/secretdb/x/secretdb/types"
)

// Handle a message to set item
// TODO: use return value of SetItem
func handleMsgSetItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetItem) (*sdk.Result, error) {
	if !k.ItemExists(ctx, msg.Filter) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, msg.Filter)
	}
	if !msg.Owner.Equals(k.GetItemOwner(ctx, msg.Filter)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.SetItem(ctx, msg.Filter, msg.Update)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
