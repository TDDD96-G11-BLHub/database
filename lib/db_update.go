package lib

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateHello() {
	fmt.Println("Hello, this is the update package")
}

// TODO: find a better solution to shifting between bson.D and []interface{}
// InsertDocument takes the client, a collection name and a bson document to be inserted.
// The client must have an open connection in the global scope for this to work.
// The function will print the unique id of the inserted document
// Keep in mind that the document must be converted to a bson D type before calling
// this function.
func InsertDocument(client mongo.Client, database string, collection string, document bson.D, docs []interface{}, fn DBFunction) {

	coll := client.Database(database).Collection(collection)

	switch fn {
	case FnInsertOne:
		res, err := coll.InsertOne(context.TODO(), document)
		if err != nil {
			panic(err)
		}
		fmt.Printf("inserted document with ID %v\n", res.InsertedID)
		return

	case FnInsertMany:
		opts := options.InsertMany().SetOrdered(false)
		res, err := coll.InsertMany(context.TODO(), docs, opts)
		if err != nil {
			panic(err)
		}
		fmt.Printf("inserted documents with IDs %v\n", res.InsertedIDs)
		return
	}

	fmt.Printf("No data was found, check collection name %s or database name %s\n", collection, database)
}

// DeleteDocument deletes the document(s) that matches the hexstring id given
// The function also takes a client, the name of the collection, the databasename and a filter
// The DBFunction enum declares which mode to be used.
// For DeleteOne the given filter must be the ObjectID to avoid deleting a random document with
// a matching value.
// Keep in mind that the client most have an open connection in the global scope for this
// to work.
// The function will print the amount of documents deleted
func DeleteDocument(client mongo.Client, database string, collection string, filter bson.D, fn DBFunction) {

	coll := client.Database(database).Collection(collection)

	//TODO: Check what this does, currently no clue
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})

	switch fn {
	case FnDeleteOne:
		res, err := coll.DeleteOne(context.TODO(), filter, opts)
		if err != nil {
			panic(err)
		}
		fmt.Printf("deleted %v documents\n", res.DeletedCount)
		return

	case FnDeleteMany:
		res, err := coll.DeleteMany(context.TODO(), filter, opts)
		if err != nil {
			panic(err)
		}
		fmt.Printf("deleted %v documents\n", res.DeletedCount)
		return
	}

	fmt.Printf("An incorrect deletemode was selected.\n Selected mode: %s\n Allowed modes: %s, %s\n", fn.String(), FnDeleteOne.String(), FnDeleteMany.String())
}
