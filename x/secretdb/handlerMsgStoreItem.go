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

// Handle a message to store item
func handleMsgStoreItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgStoreItem) (*sdk.Result, error) {
	isChild := types.IsChild

	// check sender is parent chain
	if isChild && !types.OperatorAccount.Equals(msg.GetSigners()[0]) {
		return nil, errors.New("tx from parent chain is acceptable")
	}

	var data bson.M
	err := bson.UnmarshalExtJSON([]byte(msg.Data), true, &data)
	if err != nil {
		return nil, err
	}

	var item = types.Item{
		Owner: msg.Owner,
		Data:  data,
	}
	res, err := k.StoreItem(item, !isChild)
	if err != nil {
		return nil, err
	}

	log, _ := json.Marshal(res)
	return &sdk.Result{
		Log:    fmt.Sprintf("%s", string(log)),
		Events: ctx.EventManager().Events(),
	}, nil
}
