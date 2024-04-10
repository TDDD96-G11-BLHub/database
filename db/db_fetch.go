package db

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FetchHello() {
	fmt.Println("Hello, this is the fetch package")
}

// TODO: Find a way to decode hexadecimal in error messages

// FetchOneDocument querys the database and returns one matching documents.
// An open clientconnection is required in the global scope to fetch from the database.
// Both the database and collection name must be given, as well as a query filter.
// It will return the result as marshaled jsondata
func FetchOneDocument(client mongo.Client, database string, collection string, filter bson.D) []byte {
	coll := client.Database(database).Collection(collection)
	var result_one bson.M

	err := coll.FindOne(context.TODO(), filter).Decode(&result_one)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the filter %x\n", filter)
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result_one, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)

	return jsonData
}

// FetchManyDocuments querys the database and returns all matching documents.
// An open clientconnection is required in the global scope to fetch from the database.
// Both the database and collection name must be given, as well as a query filter.
// To get all documents in the collection this function can be calle with and empty bson.D object filter
// It will return the result as marshaled jsondata
func FetchManyDocuments(client mongo.Client, database string, collection string, filter bson.D) []byte {
	coll := client.Database(database).Collection(collection)
	var result_many []bson.M

	cursor, err := coll.Find(context.TODO(), filter)

	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the filter %x\n", filter)
	}

	if err != nil {
		panic(err)
	}

	if err = cursor.All(context.TODO(), &result_many); err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result_many, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	return jsonData
}
