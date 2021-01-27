package secretdb

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptokeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authutils "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/shunail2029/SecretDB/x/secretdb/types"
)

func sendTxToMaster(msgs []sdk.Msg, cdc *codec.Codec) (sdk.TxResponse, error) {
	chainID := types.MasterChainID
	nodeURI := types.MasterURI

	// prepare CLIContext and TxBuilder
	ctx := context.CLIContext{
		FromAddress: types.OperatorAddress,
		ChainID:     chainID,
		FromName:    types.OperatorName,
	}.WithNodeURI(nodeURI)
	kb, err := cryptokeys.NewKeyring(sdk.KeyringServiceName(), types.KeyringBackend, types.CLIHome, os.Stdin)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	txBldr, err := authutils.PrepareTxBuilder(
		auth.TxBuilder{}.WithTxEncoder(authutils.GetTxEncoder(cdc)).WithKeybase(kb).WithGas(types.Gas).WithChainID(chainID),
		ctx,
	)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	// build and sign the transaction
	txBytes, err := txBldr.BuildAndSign(ctx.GetFromName(), keys.DefaultKeyPass, msgs)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	// broadcast tx to child chain
	res, err := ctx.BroadcastTxSync(txBytes)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	return res, nil
}
