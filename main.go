package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/TDDD96-G11-BLHub/dbman/peerdb"
	"github.com/google/uuid"
)

// DB PR
func main() {
	fmt.Println("Hello from BLHub database manager!")

	//C:\Users\Mattias\OneDrive\Dokument\GitHub\dbman

	//ReCreateTestDb()
	TestDb()
}

func TestDb() {
	db, err := peerdb.LoadLocalDatabase(context.TODO(), "testdb", "C:/Users/Mattias", *slog.Default())
	if err != nil {
		panic(err)
	}
	sensorcolls := db.GetCollections()
	fmt.Println(sensorcolls)

	db.Open(context.TODO(), sensorcolls[0])
	id, err := uuid.Parse("74bab0c5-1f70-4e75-a83b-3e7c60cc8970")
	if err != nil {
		panic(err)
	}

	filter := peerdb.CreateFilter("Id", "=", id)

	docs, err := peerdb.LoadCollection(context.TODO(), db, sensorcolls[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(docs)

	res := peerdb.QueryCollection(context.TODO(), db, docs, filter)

	fmt.Println(res)
}

func ReCreateTestDb() {
	RemoveDb()

	db, err := peerdb.CreateLocalDatabase(context.TODO(), "testdb", "C:/Users/Mattias", *slog.Default())
	if err != nil {
		panic(err)
	}

	sensorcolls, err := db.Create(context.TODO(), "testsensorcoll.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(sensorcolls)

	db.Open(context.TODO(), sensorcolls)

	content, err := peerdb.ReadFromCSV(context.TODO(), "C:/Users/Mattias/OneDrive/Dokument/GitHub/dbman/testdata/CONT_LOG.CSV")
	if err != nil {
		panic(err)
	}
	db.Write(context.TODO(), sensorcolls, content)
}

func RemoveDb() {
	os.Remove("C:/Users/Mattias/testdb/testsensorcoll.json")
	os.Remove("C:/Users/Mattias/testdb")
}
