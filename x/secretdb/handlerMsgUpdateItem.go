package secretdb

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/shunail2029/SecretDB/x/secretdb/keeper"
	"github.com/shunail2029/SecretDB/x/secretdb/types"
)

// Handle a message to update item
func handleMsgUpdateItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateItem) (*sdk.Result, error) {
	// if filter has "_owner", change it to msg.Owner, else add "_owner" to filter
	filter := msg.Filter
	hasOwner := false
	for idx := range filter {
		if filter[idx].Key == "_owner" {
			filter[idx].Value = msg.Owner
			hasOwner = true
			break
		}
	}
	if !hasOwner {
		filter = append(filter, bson.E{
			Key:   "_owner",
			Value: msg.Owner,
		})
	}

	if !k.ItemExists(msg.Filter) {
		filter, _ := bson.MarshalExtJSON(msg.Filter, true, false)
		return nil, fmt.Errorf("item not found with filter: %s", string(filter)) // XXX: better error might exist
	}

	res, err := k.UpdateItem(msg.Filter, msg.Update)
	if err != nil {
		return nil, err
	}

	log, _ := json.Marshal(res)
	return &sdk.Result{
		Log:    fmt.Sprintf("%s", string(log)),
		Events: ctx.EventManager().Events(),
	}, nil
}
