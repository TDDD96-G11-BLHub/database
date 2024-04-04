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

// TODO: Implement switch case for FindOne() and Find()
// Very much WIP still
func FetchCollection(client mongo.Client, name string) []byte {

	coll := client.Database("Sensordata").Collection(name)
	time := "15:40:22"
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"Time", time}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the time %s\n", time)
	}

	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", jsonData)

	return jsonData
}
