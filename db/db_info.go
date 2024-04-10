package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TestConnection sends a ping to the database via the client instance
// It panics if the connection is down, otherwise prints a confirmation
func TestConnection(client mongo.Client) {

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

}

// GetAllCollections querys the database for all the collectionnames currently in the database
// It will print out these names as well as return them as a stringarray
func GetAllCollections(client mongo.Client, database string) []string {

	result, err := client.Database(database).ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("All collections for %s: %s\n", database, result)
	return result
}

// GetAllDatabases lists all the database in the current cluster.
// It will print out these names as well as return them as a stringarray
func GetAllDatabases(client mongo.Client) []string {
	result, err := client.ListDatabaseNames(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("The cluster contains the following databases: %s\n", result)
	return result
}
