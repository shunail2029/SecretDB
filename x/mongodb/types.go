package mongodb

import "go.mongodb.org/mongo-driver/bson"

// StoreItemResult contains result of StoreItem/StoreItems
type StoreItemResult struct {
	StoredItemCount int64
}

// GetItemResult contains result of GetItem/GetItems
type GetItemResult struct {
	data bson.M
}

// SetItemResult contains result of SetItem/SetItems
type SetItemResult struct {
	MatchedCount  int64
	ModifiedCount int64
	UpsertedCount int64
}

// DeleteItemResult contains result of DeleteItem/DeleteItems
type DeleteItemResult struct {
	DeletedCount int64
}
