package Adapter_test

import (
	Adapter "adapter"
	"testing"
)

func TestAdapterPattern(t *testing.T) {
	// Create a new client.
	client := &Adapter.Client{NewData: "this is new data"}

	// Create a PostgresSQL database.
	postgresDB := &Adapter.Postgres{}

	// Save new data to PostgresSQL using the client.
	client.SaveNewData(postgresDB)

	// Create a MongoDB database.
	mongoDB := &Adapter.MongoDB{}

	// Create a MongoDB adapter.
	mongoAdapter := Adapter.NewMongoAdapter(mongoDB)

	// Save new data to MongoDB using the client and MongoDB adapter.
	client.SaveNewData(mongoAdapter)

	// Check if the MongoDB adapter has the same data.
	if data := mongoDB.GetDocument(); data != "this is new data" {
		t.Errorf("Expected MongoDB data 'Some data', but got '%v'", data)
	}
}
