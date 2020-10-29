package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	itemCollection = "items"
)

// StoreItem stores one item
func StoreItem(c *Connection, document interface{}) (*mongo.InsertOneResult, error) {
	return c.collection(itemCollection).InsertOne(c.ctx, document)
}

// StoreItems stores some items
func StoreItems(c *Connection, documents []interface{}) (*mongo.InsertManyResult, error) {
	return c.collection(itemCollection).InsertMany(c.ctx, documents)
}

// GetItem gets one item
func GetItem(c *Connection, filter interface{}) *mongo.SingleResult {
	return c.collection(itemCollection).FindOne(c.ctx, filter)
}

// GetItems gets some items
func GetItems(c *Connection, filter interface{}) (*mongo.Cursor, error) {
	return c.collection(itemCollection).Find(c.ctx, filter)
}

// SetItem updates one item
func SetItem(c *Connection, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return c.collection(itemCollection).UpdateOne(c.ctx, filter, update)
}

// SetItems updates some items
func SetItems(c *Connection, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return c.collection(itemCollection).UpdateMany(c.ctx, filter, update)
}

// DeleteItem deletes one item
func DeleteItem(c *Connection, filter interface{}) (*mongo.DeleteResult, error) {
	return c.collection(itemCollection).DeleteOne(c.ctx, filter)
}

// DeleteItems deletes some items
func DeleteItems(c *Connection, filter interface{}) (*mongo.DeleteResult, error) {
	return c.collection(itemCollection).DeleteMany(c.ctx, filter)
}
