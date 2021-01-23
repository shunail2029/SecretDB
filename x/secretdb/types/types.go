package types

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// flags
const (
	FlagIsChild       = "is-child"
	FlagParentAddress = "parent-address"
)

// parent chain params
var (
	IsChild       bool
	ParentAccount sdk.AccAddress
)

// SetParentParams ...
func SetParentParams(child bool, parentAddr string) error {
	if child && parentAddr == "" {
		return errors.New("parent address should be specified")
	}

	IsChild = child
	ParentAccount = sdk.AccAddress(parentAddr)
	return nil
}
