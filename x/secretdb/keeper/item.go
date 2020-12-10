package keeper

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shunail2029/SecretDB/x/mongodb"
	"github.com/shunail2029/SecretDB/x/secretdb/types"
	"go.mongodb.org/mongo-driver/bson"
)

// StoreItem stores a item
func (k Keeper) StoreItem(item types.Item) (mongodb.StoreItemResult, error) {
	data := insertOwner(item.Owner, item.Data)
	return mongodb.StoreItem(data)
}

// GetItem returns the item information
func (k Keeper) GetItem(iFil types.ItemFilter) (mongodb.GetItemResult, error) {
	filter := insertOwner(iFil.Owner, iFil.Filter)
	return mongodb.GetItem(filter)
}

// GetItems returns the item information
func (k Keeper) GetItems(iFil types.ItemFilter) (mongodb.GetItemResult, error) {
	filter := insertOwner(iFil.Owner, iFil.Filter)
	return mongodb.GetItems(filter)
}

// UpdateItem sets a item
func (k Keeper) UpdateItem(iFil types.ItemFilter, update bson.M) (mongodb.UpdateItemResult, error) {
	filter := insertOwner(iFil.Owner, iFil.Filter)
	return mongodb.UpdateItem(filter, update)
}

// UpdateItems sets some items
func (k Keeper) UpdateItems(iFil types.ItemFilter, update bson.M) (mongodb.UpdateItemResult, error) {
	filter := insertOwner(iFil.Owner, iFil.Filter)
	return mongodb.UpdateItems(filter, update)
}

// DeleteItem deletes a item
func (k Keeper) DeleteItem(iFil types.ItemFilter) (mongodb.DeleteItemResult, error) {
	filter := insertOwner(iFil.Owner, iFil.Filter)
	return mongodb.DeleteItem(filter)
}

// DeleteItems deletes some items
func (k Keeper) DeleteItems(iFil types.ItemFilter) (mongodb.DeleteItemResult, error) {
	filter := insertOwner(iFil.Owner, iFil.Filter)
	return mongodb.DeleteItems(filter)
}

//
// Functions used by querier
//

// getItem returns the item information
func getItem(path []string, k Keeper) ([]byte, error) {
	var filter bson.M
	err := bson.UnmarshalExtJSON([]byte(path[0]), true, &filter)
	if err != nil {
		return nil, err
	}
	_, ok := filter["_owner"]
	if !ok {
		return nil, fmt.Errorf("owner must be specified")
	}

	dbRes, err := mongodb.GetItem(filter)
	if err != nil {
		return nil, err
	}

	var res []byte
	res, err = bson.MarshalExtJSON(dbRes.Data[0], true, false)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetItems returns the item information
func getItems(path []string, k Keeper) ([]byte, error) {
	var filter bson.M
	err := bson.UnmarshalExtJSON([]byte(path[0]), true, &filter)
	if err != nil {
		return nil, err
	}
	_, ok := filter["_owner"]
	if !ok {
		return nil, fmt.Errorf("owner must be specified")
	}

	dbRes, err := mongodb.GetItems(filter)
	if err != nil {
		return nil, err
	}

	var res []byte
	for _, data := range dbRes.Data {
		res, err = bson.MarshalExtJSONAppend(res, data, true, false)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

// GetItemOwner gets owner of the item
func (k Keeper) GetItemOwner(filter bson.M) sdk.AccAddress {
	res, err := mongodb.GetItem(filter)
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
func (k Keeper) GetItemsOwner(filter bson.M) sdk.AccAddress {
	res, err := mongodb.GetItem(filter)
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
func (k Keeper) ItemExists(iFil types.ItemFilter) bool {
	filter := insertOwner(iFil.Owner, iFil.Filter)
	res, err := mongodb.GetItem(filter)
	return err == nil && res.GotItemCount > 0
}

// if filter has "_owner", change it to owner, else add "_owner" to filter
func insertOwner(owner sdk.AccAddress, m bson.M) bson.M {
	if m == nil {
		m = make(bson.M)
	}
	m["_owner"] = owner
	return m
}
