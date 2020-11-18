package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection ...
type Connection struct {
	ctx    context.Context
	cancel context.CancelFunc
	clt    *mongo.Client
	dbname string
}

// newConnection is a constructor of Connection
// TODO: enable to change URL of local MongoDB
func newConnection() *Connection {
	c := new(Connection)

	err := c.connect()
	if err != nil {
		return nil
	}
	c.dbname = "secretdb"

	return c
}

// create connection to local database
func (c Connection) connect() error {
	var err error
	c.ctx, c.cancel = context.WithTimeout(context.Background(), 20*time.Second)
	c.clt, err = mongo.Connect(c.ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	return nil
}

// disconnect closes connection to local database
func (c Connection) disconnect() {
	c.cancel()
}

// get database
func (c Connection) db() *mongo.Database {
	return c.clt.Database(c.dbname)
}

// get collection
func (c Connection) collection(name string) *mongo.Collection {
	return c.db().Collection(name)
}
