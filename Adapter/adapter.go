package Adapter

import "fmt"

// Client represents a client that interacts with a database.
type Client struct {
	NewData interface{}
}

// SaveNewData saves new data to the specified database.
func (c *Client) SaveNewData(database RelationalDatabase) {
	fmt.Printf("Process of inserting new data to the relational database was started!\n")
	database.CreateRecord(c.NewData)
	fmt.Printf("Process of inserting new data to the relational database was ended!\n")
}

// RelationalDatabase is an interface for relational databases.
type RelationalDatabase interface {
	CreateRecord(data interface{})
}

// Postgres represents a PostgreSQL database.
type Postgres struct{}

// CreateRecord creates a new record in PostgreSQL.
func (p *Postgres) CreateRecord(newData interface{}) {
	fmt.Printf("New record (%v) has been created in Postgres!\n", newData)
}

// MongoDB represents a MongoDB database.
type MongoDB struct {
	data interface{}
}

// CreateDocument creates a new document in MongoDB.
func (mdb *MongoDB) CreateDocument(newData interface{}) {
	mdb.data = newData
	fmt.Printf("New document has been created in MongoDB!\n")
}

// GetDocument retrieves the data from MongoDB.
func (mdb *MongoDB) GetDocument() interface{} {
	return mdb.data
}

// MongoAdapter is an adapter for MongoDB to implement the RelationalDatabase interface.
type MongoAdapter struct {
	mongo *MongoDB
}

// NewMongoAdapter creates a new instance of MongoAdapter.
func NewMongoAdapter(mongo *MongoDB) RelationalDatabase {
	return &MongoAdapter{mongo: mongo}
}

// CreateRecord creates a new record using MongoDB's CreateDocument method.
func (ma *MongoAdapter) CreateRecord(newData interface{}) {
	ma.mongo.CreateDocument(newData)
	fmt.Printf("New record (%v) has been created in MongoDB!\n", ma.mongo.GetDocument())
}
