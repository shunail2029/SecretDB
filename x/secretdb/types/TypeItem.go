package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Item struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
}
