package secretdb

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	mastertypes "github.com/shunail2029/SecretDB-master/x/secretdb/types"
	"github.com/shunail2029/SecretDB/x/secretdb/keeper"
	"github.com/shunail2029/SecretDB/x/secretdb/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	if !types.IsChild {
		return
	}

	if ctx.BlockHeight() == 1 {
		return
	}
	msg := mastertypes.NewMsgCreateBlockHash(types.OperatorAddress, ctx.ChainID(), ctx.BlockHeight()-1, ctx.BlockHeader().LastBlockId.Hash)
	err := msg.ValidateBasic()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	res, err := sendTxToMaster([]sdk.Msg{msg}, k.Codec())
	if err == nil {
		fmt.Println("recorded tx hash on master chain")
		fmt.Println(res)
	} else {
		fmt.Println(err.Error())
	}
}

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {}
