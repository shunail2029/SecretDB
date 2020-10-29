package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ITEM_COLLECTION = "items"
)

// store one item
func storeItem(c *MongoDBConnection, document interface{}) (*mongo.InsertOneResult, error) {
	return c.collection(ITEM_COLLECTION).InsertOne(c.ctx, document)
}

// store some items
func storeItems(c *MongoDBConnection, documents []interface{}) (*mongo.InsertManyResult, error) {
	return c.collection(ITEM_COLLECTION).InsertMany(c.ctx, documents)
}

// get one item
func getItem(c *MongoDBConnection, filter interface{}) *mongo.SingleResult {
	return c.collection(ITEM_COLLECTION).FindOne(c.ctx, filter)
}

// get some items
func getItems(c *MongoDBConnection, filter interface{}) (*mongo.Cursor, error) {
	return c.collection(ITEM_COLLECTION).Find(c.ctx, filter)
}

// delete one item
func deleteItem(c *MongoDBConnection, filter interface{}) (*mongo.DeleteResult, error) {
	return c.collection(ITEM_COLLECTION).DeleteOne(c.ctx, filter)
}

// delete some items
func deleteItems(c *MongoDBConnection, filter interface{}) (*mongo.DeleteResult, error) {
	return c.collection(ITEM_COLLECTION).DeleteMany(c.ctx, filter)
}
