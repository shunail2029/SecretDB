package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateItem{}

type MsgCreateItem struct {
	ID      string
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgCreateItem(creator sdk.AccAddress) MsgCreateItem {
	return MsgCreateItem{
		ID:      uuid.New().String(),
		Creator: creator,
	}
}

func (msg MsgCreateItem) Route() string {
	return RouterKey
}

func (msg MsgCreateItem) Type() string {
	return "CreateItem"
}

func (msg MsgCreateItem) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateItem) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
