package keeper

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shunail2029/SecretDB/x/mongodb"
	"github.com/shunail2029/SecretDB/x/secretdb/types"
	"go.mongodb.org/mongo-driver/bson"
)

// StoreItem stores a item
func (k Keeper) StoreItem(item types.Item) (mongodb.StoreItemResult, error) {
	data := item.Data
	data["_owner"] = item.Owner
	return mongodb.StoreItem(k.Conn, data)
}

// GetItem returns the item information
func (k Keeper) GetItem(filter bson.D) (mongodb.GetItemResult, error) {
	return mongodb.GetItem(k.Conn, filter)
}

// GetItems returns the item information
func (k Keeper) GetItems(filter bson.D) (mongodb.GetItemResult, error) {
	return mongodb.GetItems(k.Conn, filter)
}

// UpdateItem sets a item
func (k Keeper) UpdateItem(filter bson.D, update bson.D) (mongodb.UpdateItemResult, error) {
	return mongodb.UpdateItem(k.Conn, filter, update)
}

// UpdateItems sets some items
func (k Keeper) UpdateItems(filter bson.D, update bson.D) (mongodb.UpdateItemResult, error) {
	return mongodb.UpdateItems(k.Conn, filter, update)
}

// DeleteItem deletes a item
func (k Keeper) DeleteItem(filter bson.D) (mongodb.DeleteItemResult, error) {
	return mongodb.DeleteItem(k.Conn, filter)
}

// DeleteItems deletes some items
func (k Keeper) DeleteItems(filter bson.D) (mongodb.DeleteItemResult, error) {
	return mongodb.DeleteItems(k.Conn, filter)
}

//
// Functions used by querier
//

// getItem returns the item information
func getItem(path []string, k Keeper) ([]byte, error) {
	var filter bson.D
	err := bson.Unmarshal([]byte(path[0]), &filter)
	if err != nil {
		return nil, err
	}
	dbRes, err := mongodb.GetItem(k.Conn, filter)
	if err != nil {
		return nil, err
	}

	var res []byte
	res, err = bson.Marshal(dbRes.Data[0])
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetItems returns the item information
func getItems(path []string, k Keeper) ([]byte, error) {
	var filter bson.D
	err := bson.Unmarshal([]byte(path[0]), &filter)
	if err != nil {
		return nil, err
	}
	dbRes, err := mongodb.GetItems(k.Conn, filter)
	if err != nil {
		return nil, err
	}

	var res []byte
	for _, data := range dbRes.Data {
		res, err = bson.MarshalAppend(res, data)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

// GetItemOwner gets owner of the item
func (k Keeper) GetItemOwner(filter bson.D) sdk.AccAddress {
	res, err := mongodb.GetItem(k.Conn, filter)
	if err != nil || res.GotItemCount != 1 {
		return nil
	}
	switch addr := res.Data[0]["_owner"].(type) {
	case sdk.AccAddress:
		return addr
	default:
		return nil
	}
}

// GetItemsOwner gets owner of the items
// If one owner owns all items, return address of the owner
func (k Keeper) GetItemsOwner(filter bson.D) sdk.AccAddress {
	res, err := mongodb.GetItem(k.Conn, filter)
	if err != nil || res.GotItemCount == 0 {
		return nil
	}
	switch addr := res.Data[0]["_owner"].(type) { // type assertion of res.Data[0]["_owner"]
	case sdk.AccAddress:
		for _, data := range res.Data {
			switch dataAddr := data["_owner"].(type) { // type assertion of data["_owner"]
			case sdk.AccAddress:
				if !bytes.Equal(dataAddr, addr) {
					return nil
				}
			default:
				return nil
			}
		}
		return addr
	default:
		return nil
	}
}

// ItemExists checks if the key exists in the store
func (k Keeper) ItemExists(filter bson.D) bool {
	res, err := mongodb.GetItem(k.Conn, filter)
	return err == nil && res.GotItemCount > 0
}
