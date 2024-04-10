package db

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

//TODO: Update comments

// InsertOneDocument takes the client, a collection name and a bson document to be inserted.
// The client must have an open connection in the global scope for this to work.
// The function will print the unique id of the inserted document
// Keep in mind that the document must be converted to a bson D type before calling
// this function.
func InsertOneDocument(client mongo.Client, database string, collection string, document bson.D) {

	coll := client.Database(database).Collection(collection)

	res, err := coll.InsertOne(context.TODO(), document)
	if err != nil {
		panic(err)
	}
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)
}

// InsertManyDocuments takes the client, a collection name and an interface with bson documents to be inserted.
// The client must have an open connection in the global scope for this to work.
// The function will print the unique ids of all the inserted documents
// Keep in mind that the documents must be converted to a bson D type and put into an interface before calling
// this function.
func InsertManyDocuments(client mongo.Client, database string, collection string, docs []interface{}) {

	coll := client.Database(database).Collection(collection)

	opts := options.InsertMany().SetOrdered(false)
	res, err := coll.InsertMany(context.TODO(), docs, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("inserted documents with IDs %v\n", res.InsertedIDs)
}

// DeleteOneDocument deletes the document that matches the hexstring id given
// The function also takes a client, the name of the collection, the databasename
// The id should be a bson.D object of the form bson.D{{"_id", <hexid>}}
// Keep in mind that the client most have an open connection in the global scope for this
// to work.
// The function will print the amount of documents deleted
func DeleteOneDocument(client mongo.Client, database string, collection string, id bson.D) {

	coll := client.Database(database).Collection(collection)

	//TODO: Check what this does, currently no clue. Sets some sort of rules for query?
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})

	res, err := coll.DeleteOne(context.TODO(), id, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
}

// DeleteManyDocuments deletes all the documents that matches the filter given
// The function also takes a client, the name of the collection, the databasename
// Keep in mind that the client most have an open connection in the global scope for this
// to work.
// To delete all documents it can be called with an empty bson.D object filter although it would be
// better to drop the collection
// The function will print the amount of documents deleted
func DeleteManyDocuments(client mongo.Client, database string, collection string, filter bson.D) {

	coll := client.Database(database).Collection(collection)

	//TODO: Check what this does, currently no clue. Sets some sort of rules for query?
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})

	res, err := coll.DeleteMany(context.TODO(), filter, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
}

// NewCollection creates a new collection on the given database with the given name.
// This function currently uses no options when creating the collection which might be unsafe
// and should probably be implemented before deployment.
func NewCollection(client mongo.Client, database string, collection string) {

	err := client.Database(database).CreateCollection(context.TODO(), collection)
	if err != nil {
		panic(err)
	}
}

// DropDatabase drops the entire database from the cluster.
// This function ignores the namspace not found error so it won't crash
// the program if the database doesn't exist.
// Keep in mind that this function is irreversible so be careful when using.
func DropDatabase(client mongo.Client, database string) {

	err := client.Database(database).Drop(context.TODO())
	if err != nil {
		panic(err)
	}
}

// DropCollection drops the entire collection from the database.
// This function ignores the namspace not found error so it won't crash
// the program if the collection doesn't exist.
// Keep in mind that this function is irreversible so be careful when using.
func DropCollection(client mongo.Client, database string, collection string) {

	err := client.Database(database).Collection(collection).Drop(context.TODO())
	if err != nil {
		panic(err)
	}
}
