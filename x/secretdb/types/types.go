package types

import (
	"errors"
	"os"

	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

// flags
const (
	FlagIsChild        = "is-child"
	FlagOperatorName   = "operator-name"
	FlagKeyringBackend = "keyring-backend"
	FlagCLIHome        = "cli-home" // to use keyring
	FlagGas            = "gas"
	FlagMasterURI      = "master-uri"
	FlagMasterChainID  = "master-chainid"
)

// parent chain params
var (
	IsChild         bool
	OperatorName    string
	OperatorAddress sdk.AccAddress
	OperatorPubkey  crypto.PubKey
	KeyringBackend  string
	CLIHome         string
	Gas             uint64
	MasterURI       string
	MasterChainID   string
)

// SetParams ...
func SetParams(child bool, name, keyringBackend, cliHome string, gas uint64, uri, chainID string) error {
	IsChild = child

	if name == "" {
		return errors.New("operator name must be specified")
	}
	OperatorName = name

	if keyringBackend == "" {
		return errors.New("keyring backend must be specified")
	}
	KeyringBackend = keyringBackend

	CLIHome = cliHome
	Gas = gas

	kb, err := keys.NewKeyring(sdk.KeyringServiceName(), KeyringBackend, CLIHome, os.Stdin)
	if err != nil {
		return err
	}
	info, err := kb.Get(OperatorName)
	if err != nil {
		return err
	}
	OperatorAddress = info.GetAddress()
	OperatorPubkey = info.GetPubKey()

	MasterURI = uri
	MasterChainID = chainID
	return nil
}
