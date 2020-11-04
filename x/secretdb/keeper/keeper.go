package keeper

import (
	"fmt"

	"github.com/shunail2029/SecretDB/x/mongodb"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/shunail2029/SecretDB/x/secretdb/types"
)

// Keeper of the secretdb store
type Keeper struct {
	CoinKeeper bank.Keeper
	Conn       *mongodb.Connection
	cdc        *codec.Codec
}

// NewKeeper creates a secretdb keeper
func NewKeeper(coinKeeper bank.Keeper, cdc *codec.Codec, conn *mongodb.Connection) Keeper {
	keeper := Keeper{
		CoinKeeper: coinKeeper,
		Conn:       conn,
		cdc:        cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
