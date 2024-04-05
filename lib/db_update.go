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

func InsertDocument(client mongo.Client, name string, document bson.D) {

	coll := client.Database("Sensordata").Collection(name)

	res, err := coll.InsertOne(context.TODO(), document)
	if err != nil {
		panic(err)
	}
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)
}

func DeleteDocument(client mongo.Client, name string, id_hex string) {

	coll := client.Database("Sensordata").Collection(name)
	id, err := primitive.ObjectIDFromHex(id_hex)
	if err != nil {
		panic(err)
	}
	filter := bson.D{{"_id", id}}
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
