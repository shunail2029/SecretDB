package secretdb

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/shunail2029/secretdb/x/secretdb/keeper"
	"github.com/shunail2029/secretdb/x/secretdb/types"
)

// NewHandler ...
// TODO: add MsgCreateItems to create some items at once
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case types.MsgCreateItem:
			return handleMsgCreateItem(ctx, k, msg)
		case types.MsgSetItem:
			return handleMsgSetItem(ctx, k, msg)
		case types.MsgSetItems:
			return handleMsgSetItems(ctx, k, msg)
		case types.MsgDeleteItem:
			return handleMsgDeleteItem(ctx, k, msg)
		case types.MsgDeleteItems:
			return handleMsgDeleteItems(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
