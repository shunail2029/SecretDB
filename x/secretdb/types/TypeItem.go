package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"go.mongodb.org/mongo-driver/bson"
)

// Item is a type of data stored in MongoDB
type Item struct {
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
	Data  bson.M         `json:"data" yaml:"data"`
}
