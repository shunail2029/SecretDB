package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"go.mongodb.org/mongo-driver/bson"
)

var _ sdk.Msg = &MsgSetItem{}

// MsgSetItem is message type to set item
type MsgSetItem struct {
	Owner  sdk.AccAddress `json:"owner" yaml:"owner"`
	Filter bson.D         `json:"filter" yaml:"filter"`
	Update bson.D         `json:"update" yaml:"update"`
}

// NewMsgSetItem returns new MsgSetItem
func NewMsgSetItem(owner sdk.AccAddress, filter bson.D, update bson.D) MsgSetItem {
	return MsgSetItem{
		Owner:  owner,
		Filter: filter,
		Update: update,
	}
}

// Route ...
func (msg MsgSetItem) Route() string {
	return RouterKey
}

// Type ...
func (msg MsgSetItem) Type() string {
	return "SetItem"
}

// GetSigners ...
func (msg MsgSetItem) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

// GetSignBytes ...
func (msg MsgSetItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg MsgSetItem) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
	}
	return nil
}
