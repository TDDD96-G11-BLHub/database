package lib

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateHello() {
	fmt.Println("Hello, this is the update package")
}

// InsertDocument takes the client, a collection name and a bson document to be inserted.
// The client must have an open connection in the global scope for this to work.
// The function will print the unique id of the inserted document
// Keep in mind that the document must be converted to a bson D type before calling
// this function.
func InsertDocument(client mongo.Client, name string, document bson.D, fn DBFunction) {

	coll := client.Database("Sensordata").Collection(name)

	res, err := coll.InsertOne(context.TODO(), document)
	if err != nil {
		panic(err)
	}
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)
}

// DeleteDocument deletes the document that matches the hexstring id given
// The function also takes a client and the name of the collection to be updated
// Keep in mind that the client most have an open connection in the global scope for this
// to work.
// The function will print the amount of documents deleted (currently 1 or 0)
func DeleteDocument(client mongo.Client, name string, id_hex string, fn DBFunction) {

	coll := client.Database("Sensordata").Collection(name)
	id, err := primitive.ObjectIDFromHex(id_hex)
	if err != nil {
		panic(err)
	}
	filter := bson.D{{"_id", id}}
	//TODO: Check what this does, currently no clue
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})

	res, err := coll.DeleteOne(context.TODO(), filter, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
}
