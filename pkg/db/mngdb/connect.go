// Package mngdb provides a centralized interface for managing all MongoDB operations
// within the project. It handles establishing connections, managing the MongoDB client,
// and encapsulating common database tasks such as CRUD operations.
//
// The package uses the official MongoDB Go driver and relies on the `config` package
// to retrieve configuration settings such as the MongoDB URI.
package mngdb

import (
	"context"
	"log"

	"github.com/schedlinkk/server/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

// Connect establishes a connection to the MongoDB server and initializes the client.
// The MongoDB URI is retrieved from the `config` package. This function ensures that
// the database server is reachable by performing a ping test.
//
// Returns an error if:
//   - The connection to the MongoDB server cannot be established.
//   - The server cannot be pinged successfully.
//
// Example:
//
//	if err := mngdb.Connect(); err != nil {
//	    log.Fatalf("Failed to connect to MongoDB: %v", err)
//	}
//	defer func() {
//	    if err := mngdb.Disconnect(); err != nil {
//	        log.Printf("Error disconnecting from MongoDB: %v", err)
//	    }
//	}()
func Connect() error {
	uri := config.MongoURI()

	log.Println("Start connecting to MongoDB...")
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	log.Println("Pinging MongoDB...")
	if err := c.Ping(context.TODO(), readpref.Primary()); err != nil {
		return err
	}

	client = c

	log.Println("Successfully connected to MongoDB.")

	return nil

}

// Disconnect gracefully terminates the connection to the MongoDB server
// and releases all resources used by the client.
//
// Returns an error if the disconnection process fails.
//
// Example:
//
//	defer func() {
//	    if err := mngdb.Disconnect(); err != nil {
//	        log.Printf("Error disconnecting from MongoDB: %v", err)
//	    }
//	}()
func Disconnect() error {
	if err := client.Disconnect(context.TODO()); err != nil {
		return err
	}

	return nil
}
