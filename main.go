package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/TDDD96-G11-BLHub/dbman/peerdb"
)

// DB PR
func main() {
	fmt.Println("Hello from BLHub database manager!")

	//C:\Users\Mattias\OneDrive\Dokument\GitHub\dbman

	// db, err := peerdb.CreateLocalDatabase(context.TODO(), "testdb", "C:/Users/Mattias", *slog.Default())
	// if err != nil {
	// 	panic(err)
	// }

	// sensorcolls, err := db.Create(context.TODO(), "testsensorcoll.json")
	// if err != nil {
	// 	panic(err)
	// }
	db := peerdb.LoadLocalDatabase(context.TODO(), "testdb", "C:/Users/Mattias", *slog.Default())
	sensorcolls := db.GetCollections()

	db.Open(context.TODO(), sensorcolls[0])
	content, err := peerdb.ReadFromCSV(context.TODO(), "C:/Users/Mattias/OneDrive/Dokument/GitHub/dbman/testdata/CONT_LOG.CSV")
	if err != nil {
		panic(err)
	}
	db.Write(context.TODO(), sensorcolls[0], content)

	filter := peerdb.CreateFilter("Roll", "=", 0.438369)

	docs, err := peerdb.LoadCollection(context.TODO(), db, sensorcolls[0])
	if err != nil {
		panic(err)
	}

	res := peerdb.QueryCollection(context.TODO(), db, docs, filter)

	fmt.Println(res)
}
