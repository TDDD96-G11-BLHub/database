package lib

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectHello() {
	fmt.Println("Hello, this is the connect package")
}

// SetupConnection creates a connection to the mongodb database given the uri setup in the environment variable
// If an error would occur in the .env file or while loading the variable the function will log the error
// If an error would occur with the client connection, the function will close the connection and panic
// The returned value is a Client struct which can be used to handle the connection to the database
func SetupConnection() mongo.Client {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	//uri := "mongodb+srv://mattkarl2001:Homer2001@blhub-test.1hxmrul.mongodb.net/"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	return *client
}
