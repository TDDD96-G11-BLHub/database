package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/TDDD96-G11-BLHub/dbman/lib"
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

	lib.TestConnection(*client)
	lib.FetchCollection(*client, "deepoidsensor")
	lib.ConnectHello()
	lib.FetchHello()
	lib.UpdateHello()

}
