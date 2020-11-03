package secretdb

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/shunail2029/secretdb/x/secretdb/keeper"
	"github.com/shunail2029/secretdb/x/secretdb/types"
)

// Handle a message to delete item
// TODO: use return value of DeleteItem
func handleMsgDeleteItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteItem) (*sdk.Result, error) {
	if !k.ItemExists(ctx, msg.Filter) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, msg.Filter)
	}
	if !msg.Owner.Equals(k.GetItemOwner(ctx, msg.Filter)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteItem(ctx, msg.Filter)
	return &sdk.Result{}, nil
}
