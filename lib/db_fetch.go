package lib

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
// FetchDocument querys the database and returns matching documents.
// An open clientconnection is required in the global scope to fetch from the database.
// Both the database and collection name must be given, as well as a query filter.
// The fetch mode can be set by choosing the function enum
// It will return the result as marshaled jsondata
func FetchDocument(client mongo.Client, database string, collection string, filter bson.D, fn DBFunction) []byte {
	coll := client.Database(database).Collection(collection)
	var result_one bson.M
	var result_many []bson.M

	switch fn {
	case FnFindOne:
		fmt.Println(FnFindOne.String())
		err := coll.FindOne(context.TODO(), filter).Decode(&result_one)

		if err == mongo.ErrNoDocuments {
			fmt.Printf("No document was found with the filter %x\n", filter)
		}

		if err != nil {
			panic(err)
		}

	case FnFindMany:
		fmt.Println(FnFindMany.String())
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
	}

	if result_one != nil {
		jsonData, err := json.MarshalIndent(result_one, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", jsonData)

		return jsonData
	}

	if result_many != nil {
		jsonData, err := json.MarshalIndent(result_many, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", jsonData)
		return jsonData
	}
	fmt.Printf("No data was found, check query filter %x, collection name %s or database name %s", filter, collection, database)
	return nil
}
