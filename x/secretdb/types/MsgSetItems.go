package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"go.mongodb.org/mongo-driver/bson"
)

var _ sdk.Msg = &MsgSetItems{}

// MsgSetItems is message type to set some items
type MsgSetItems struct {
	Owner  sdk.AccAddress `json:"owner" yaml:"owner"`
	Filter bson.D         `json:"filter" yaml:"filter"`
	Update bson.D         `json:"update" yaml:"update"`
}

// NewMsgSetItems returns new MsgSetItems
func NewMsgSetItems(owner sdk.AccAddress, filter bson.D, update bson.D) MsgSetItems {
	return MsgSetItems{
		Owner:  owner,
		Filter: filter,
		Update: update,
	}
}

// Route ...
func (msg MsgSetItems) Route() string {
	return RouterKey
}

// Type ...
func (msg MsgSetItems) Type() string {
	return "SetItems"
}

// GetSigners ...
func (msg MsgSetItems) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

// GetSignBytes ...
func (msg MsgSetItems) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg MsgSetItems) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
	}
	return nil
}
