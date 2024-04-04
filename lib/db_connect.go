package lib

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConnectHello() {
	fmt.Println("Hello, this is the connect package")
}

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
