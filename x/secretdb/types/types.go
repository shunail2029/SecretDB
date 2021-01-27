package types

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// flags
const (
	FlagIsChild         = "is-child"
	FlagOperatorName    = "operator-name"
	FlagOperatorAddress = "operator-address"
	FlagKeyringBackend  = "keyring-backend"
	FlagCLIHome         = "cli-home" // to use keyring
	FlagGas             = "gas"
	FlagMasterURI       = "master-uri"
	FlagMasterChainID   = "master-chainid"
)

// parent chain params
var (
	IsChild         bool
	OperatorName    string
	OperatorAddress sdk.AccAddress
	KeyringBackend  string
	CLIHome         string
	Gas             uint64
	MasterURI       string
	MasterChainID   string
)

// SetParams ...
func SetParams(child bool, name, address, keyringBackend, cliHome string, gas uint64, uri, chainID string) error {
	var err error

	IsChild = child
	if !child {
		return nil
	}

	if name == "" {
		return errors.New("operator name must be specified")
	}
	OperatorName = name

	if address == "" {
		return errors.New("operator address must be specified")
	}
	OperatorAddress, err = sdk.AccAddressFromBech32(address)
	if err != nil {
		return err
	}

	if keyringBackend == "" {
		return errors.New("keyring backend must be specified")
	}
	KeyringBackend = keyringBackend

	CLIHome = cliHome
	Gas = gas

	MasterURI = uri
	MasterChainID = chainID
	return nil
}
