package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection ...
type Connection struct {
	ctx    context.Context
	clt    *mongo.Client
	dbname string
}

// NewConnection is a constructor of Connection
// TODO: enable to change URL of local MongoDB
func NewConnection() *Connection {
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
	c.clt, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	c.ctx = context.Background()
	err = c.clt.Connect(c.ctx)
	if err != nil {
		return err
	}
	return nil
}

// get database
func (c Connection) db() *mongo.Database {
	return c.clt.Database(c.dbname)
}

// get collection
func (c Connection) collection(name string) *mongo.Collection {
	return c.db().Collection(name)
}
