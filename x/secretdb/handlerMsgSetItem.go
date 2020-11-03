package secretdb

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/shunail2029/secretdb/x/secretdb/keeper"
	"github.com/shunail2029/secretdb/x/secretdb/types"
)

// Handle a message to set item
func handleMsgSetItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetItem) (*sdk.Result, error) {
	if !k.ItemExists(msg.Filter) {
		filter, _ := bson.Marshal(msg.Filter)
		return nil, fmt.Errorf("item not found with filter: %s", string(filter)) // XXX: better error might exist
	}
	if !msg.Owner.Equals(k.GetItemOwner(msg.Filter)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	res, err := k.SetItem(msg.Filter, msg.Update)
	if err != nil {
		return nil, err
	}

	log, _ := json.Marshal(res)
	return &sdk.Result{
		Log:    fmt.Sprintf("%s", string(log)),
		Events: ctx.EventManager().Events(),
	}, nil
}
