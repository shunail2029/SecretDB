package keeper

import (
	// this line is used by starport scaffolding # 1
	"github.com/shunail2029/SecretDB/x/secretdb/types"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for secretdb clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryGetOperatorPubkey:
			return getOperatorPubkey(k)
		case types.QueryGetItem:
			return getItem(path[1:], k, !types.IsChild)
		case types.QueryGetItems:
			return getItems(path[1:], k, !types.IsChild)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown secretdb query endpoint")
		}
	}
}
