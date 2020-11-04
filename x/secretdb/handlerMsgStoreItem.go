package secretdb

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/shunail2029/SecretDB/x/secretdb/keeper"
	"github.com/shunail2029/SecretDB/x/secretdb/types"
)

// Handle a message to store item
func handleMsgStoreItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgStoreItem) (*sdk.Result, error) {
	var item = types.Item{
		Owner: msg.Owner,
		Data:  msg.Data,
	}
	res, err := k.StoreItem(item)
	if err != nil {
		return nil, err
	}

	log, _ := json.Marshal(res)
	return &sdk.Result{
		Log:    fmt.Sprintf("%s", string(log)),
		Events: ctx.EventManager().Events(),
	}, nil
}
