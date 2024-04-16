package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB PR
func main() {
	fmt.Println("Hello from BLHub database manager!")

	// This section creates a connection to the mongodb database given the uri setup in the environment variable
	// If an error would occur in the .env file or while loading the variable the function will log the error
	// If an error would occur with the client connection, the function will close the connection and panic
	// The best practice is for the Client instance to be in the global scope which is why it is in main

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	//How to create ObjectID values
	//id, err := primitive.ObjectIDFromHex("6613bc82b3d073c74f4d038a")
	//if err != nil {
	//	panic(err)
	//}

	//Example of multiple docs
	//docs := []interface{}{
	//	bson.D{{"Time", "14:52:22"},
	//		{"Roll", 0.723491},
	//		{"Pitch", -3.248201},
	//		{"Yaw", 0.345234}},
	//	bson.D{{"Time", "14:52:22"},
	//		{"Roll", 0.723291},
	//		{"Pitch", -3.238201},
	//		{"Yaw", 0.345214}},
	//}

	// Example of a document
	// document := bson.D{
	//	{"Time", "14:52:22"},
	//	{"Roll", 0.723491},
	//	{"Pitch", -3.248201},
	//	{"Yaw", 0.345234}}

	//Example of a filter
	//filter := bson.D{{"Time", "15:40:22"}}
	//
	//db.TestConnection(client)
	//db.GetAllDatabases(client)
	//db.GetAllCollections(client, "Sensordata")
	//db.FetchOneDocument(client, "Sensordata", "deepoidsensor", filter)
	//db.FetchManyDocuments(client, "Sensordata", "deepoidsensor", filter)
	// db.NewCollection(client, "Sensordata", "sensortyp2")
	// db.InsertManyDocuments(client, "Sensordata", "deepoidsensor", docs)
	// db.DeleteManyDocuments(client, "Sensordata", "deepoidsensor", filter)
}
