package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConnection struct {
	ctx    context.Context
	clt    *mongo.Client
	dbname string
}

// constructor
func NewMongoDBConnection() *MongoDBConnection {
	c := new(MongoDBConnection)

	err := c.connect()
	if err != nil {
		return nil
	}
	c.dbname = "secretdb"

	return c
}

// create connection to local database
func (c MongoDBConnection) connect() error {
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
func (c MongoDBConnection) db() *mongo.Database {
	return c.clt.Database(c.dbname)
}

// get collection
func (c MongoDBConnection) collection(name string) *mongo.Collection {
	return c.db().Collection(name)
}
