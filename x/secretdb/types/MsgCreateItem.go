package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"go.mongodb.org/mongo-driver/bson"
)

var _ sdk.Msg = &MsgCreateItem{}

// MsgCreateItem is a message type to create item
type MsgCreateItem struct {
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
	Data  bson.M         `json:"data" yaml:"data"`
}

// NewMsgCreateItem returns new MsgCreateItem
func NewMsgCreateItem(owner sdk.AccAddress, data bson.M) MsgCreateItem {
	return MsgCreateItem{
		Owner: owner,
		Data:  data,
	}
}

// Route ...
func (msg MsgCreateItem) Route() string {
	return RouterKey
}

// Type ...
func (msg MsgCreateItem) Type() string {
	return "CreateItem"
}

// GetSigners ...
func (msg MsgCreateItem) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

// GetSignBytes ...
func (msg MsgCreateItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg MsgCreateItem) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
	}
	return nil
}
