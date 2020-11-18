package mongodb

import "go.mongodb.org/mongo-driver/bson"

const (
	itemCollection = "items"
)

// StoreItem stores one item
func StoreItem(document interface{}) (StoreItemResult, error) {
	// create connection to database
	c := newConnection()
	defer c.disconnect()

	_, err := c.collection(itemCollection).InsertOne(c.ctx, document)
	if err != nil {
		return StoreItemResult{}, err
	}
	return StoreItemResult{
		StoredItemCount: 1,
	}, nil
}

// StoreItems stores some items
func StoreItems(documents []interface{}) (StoreItemResult, error) {
	// create connection to database
	c := newConnection()
	defer c.disconnect()

	res, err := c.collection(itemCollection).InsertMany(c.ctx, documents)
	if err != nil {
		return StoreItemResult{}, err
	}
	return StoreItemResult{
		StoredItemCount: int64(len(res.InsertedIDs)),
	}, nil
}

// GetItem gets one item
func GetItem(filter interface{}) (GetItemResult, error) {
	// create connection to database
	c := newConnection()
	defer c.disconnect()

	var res bson.M
	err := c.collection(itemCollection).FindOne(c.ctx, filter).Decode(&res)
	if err != nil {
		return GetItemResult{}, err
	}
	return GetItemResult{
		GotItemCount: 1,
		Data:         []bson.M{res},
	}, nil
}

// GetItems gets some items
func GetItems(filter interface{}) (GetItemResult, error) {
	// create connection to database
	c := newConnection()
	defer c.disconnect()

	cursor, err := c.collection(itemCollection).Find(c.ctx, filter)
	if err != nil {
		return GetItemResult{}, err
	}

	var res []bson.M
	if err = cursor.All(c.ctx, &res); err != nil {
		return GetItemResult{}, err
	}
	return GetItemResult{
		GotItemCount: int64(len(res)),
		Data:         res,
	}, nil
}

// UpdateItem updates one item
func UpdateItem(filter interface{}, update interface{}) (UpdateItemResult, error) {
	// create connection to database
	c := newConnection()
	defer c.disconnect()

	res, err := c.collection(itemCollection).UpdateOne(c.ctx, filter, update)
	if err != nil {
		return UpdateItemResult{}, err
	}
	return UpdateItemResult{
		MatchedCount:  res.MatchedCount,
		ModifiedCount: res.ModifiedCount,
		UpsertedCount: res.UpsertedCount,
	}, nil
}

// UpdateItems updates some items
func UpdateItems(filter interface{}, update interface{}) (UpdateItemResult, error) {
	// create connection to database
	c := newConnection()
	defer c.disconnect()

	res, err := c.collection(itemCollection).UpdateMany(c.ctx, filter, update)
	if err != nil {
		return UpdateItemResult{}, err
	}
	return UpdateItemResult{
		MatchedCount:  res.MatchedCount,
		ModifiedCount: res.ModifiedCount,
		UpsertedCount: res.UpsertedCount,
	}, nil
}

// DeleteItem deletes one item
func DeleteItem(filter interface{}) (DeleteItemResult, error) {
	// create connection to database
	c := newConnection()
	defer c.disconnect()

	res, err := c.collection(itemCollection).DeleteOne(c.ctx, filter)
	if err != nil {
		return DeleteItemResult{}, err
	}
	return DeleteItemResult{
		DeletedCount: res.DeletedCount,
	}, nil
}

// DeleteItems deletes some items
func DeleteItems(filter interface{}) (DeleteItemResult, error) {
	// create connection to database
	c := newConnection()
	defer c.disconnect()

	res, err := c.collection(itemCollection).DeleteMany(c.ctx, filter)
	if err != nil {
		return DeleteItemResult{}, err
	}
	return DeleteItemResult{
		DeletedCount: res.DeletedCount,
	}, nil
}
