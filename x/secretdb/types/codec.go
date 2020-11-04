package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// TODO: check what codec does and how to use codec

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
	cdc.RegisterConcrete(MsgStoreItem{}, "secretdb/StoreItem", nil)
	cdc.RegisterConcrete(MsgUpdateItem{}, "secretdb/UpdateItem", nil)
	cdc.RegisterConcrete(MsgUpdateItems{}, "secretdb/UpdateItems", nil)
	cdc.RegisterConcrete(MsgDeleteItem{}, "secretdb/DeleteItem", nil)
	cdc.RegisterConcrete(MsgDeleteItems{}, "secretdb/DeleteItems", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
