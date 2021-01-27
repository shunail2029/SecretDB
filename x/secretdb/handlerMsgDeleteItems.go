package secretdb

import (
	"encoding/json"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/shunail2029/SecretDB/x/secretdb/keeper"
	"github.com/shunail2029/SecretDB/x/secretdb/types"
)

// Handle a message to delete some items
func handleMsgDeleteItems(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteItems) (*sdk.Result, error) {
	isChild := types.IsChild

	// check sender is parent chain
	if isChild && !types.OperatorAccount.Equals(msg.GetSigners()[0]) {
		return nil, errors.New("tx from parent chain is acceptable")
	}

	var filter bson.M
	err := bson.UnmarshalExtJSON([]byte(msg.Filter), true, &filter)
	if err != nil {
		return nil, err
	}

	iFil := types.ItemFilter{
		Owner:  msg.Owner,
		Filter: filter,
	}

	if !k.ItemExists(iFil, !isChild) {
		filter, _ := bson.MarshalExtJSON(iFil.Filter, true, false)
		return nil, fmt.Errorf("item not found with filter: %s", string(filter)) // XXX: better error might exist
	}

	res, err := k.DeleteItems(iFil, !isChild)
	if err != nil {
		return nil, err
	}

	log, _ := json.Marshal(res)
	return &sdk.Result{
		Log: fmt.Sprintf("%s", string(log)),
	}, nil
}
