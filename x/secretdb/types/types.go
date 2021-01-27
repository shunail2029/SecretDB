package types

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// flags
const (
	FlagIsChild         = "is-child"
	FlagOperatorAddress = "operator-address"
)

// parent chain params
var (
	IsChild         bool
	OperatorAccount sdk.AccAddress
)

// SetParams ...
func SetParams(child bool, operatorAddr string) error {
	var err error

	if child && operatorAddr == "" {
		return errors.New("operator address should be specified")
	}

	IsChild = child
	OperatorAccount, err = sdk.AccAddressFromBech32(operatorAddr)
	if err != nil {
		return err
	}

	return nil
}
