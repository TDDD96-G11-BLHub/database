package peerdb

import (
	"context"
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/google/uuid"
)

//TODO: Create different types of LocalDB struct. (Document, user, sensor, more...)
//TODO: Use correct type of context in functions, not just TODO
//TODO: Create and use slog in functions
//TODO: More descriptive errors, maybe even just return them
//TODO: Maybe just use string instead of address package
//TODO: Add Defer functions to close files on crash
//TODO: Add keys to collections in write function
//TODO: Add loading collections to LoadLocalDatabase

type LocalDBBase interface {
	Open(ctx context.Context, collection *Collection) (*os.File, error)

	Create(ctx context.Context, name string) (Collection, error)

	Write(ctx context.Context, collection *Collection, content []byte) (int, error)

	GetAddress(ctx context.Context, name string) (string, error)

	Logger() *slog.Logger

	GetCollections() []*Collection
}

type LocalDB struct {
	name        string
	path        string
	collections []*Collection
	log         *slog.Logger
}

type Collection struct {
	name string
	file *os.File
	keys []byte
	log  *slog.Logger
}

// type CreateCollOptions struct {
// 	Directory   string
// 	LocalOnly   bool
// 	Overwrite   bool
// 	ContentType string
// 	Logger      *slog.Logger
// }

func (db *LocalDB) Create(ctx context.Context, name string) (*Collection, error) {
	file, err := os.Create(db.path + "/" + name)
	if err != nil {
		db.log.Error("Could not create file")
	}
	file.Close()
	coll := Collection{name: name, file: nil, keys: []byte(""), log: db.log}
	db.collections = append(db.collections, &coll)
	return &coll, err
}

func CreateLocalDatabase(ctx context.Context, name string, path string, log slog.Logger) (*LocalDB, error) {
	err := os.Mkdir(path+"/"+name, 0777)
	if err != nil {
		log.Error("Could not create directory")
	}
	var collections []*Collection
	db := LocalDB{name: name, path: path + "/" + name, collections: collections, log: &log}
	return &db, err
}

func LoadLocalDatabase(ctx context.Context, name string, path string, log slog.Logger) (*LocalDB, error) {
	var collections []*Collection
	db := LocalDB{name: name, path: path + "/" + name, collections: collections, log: &log}
	files, err := os.ReadDir(path + "/" + name)
	if err != nil {
		log.Error("Could not load directory")
		fmt.Println(err)
	}
	var i int
	for i = 0; i < len(files); i++ {
		coll := Collection{files[i].Name(), nil, []byte(" "), &log}
		db.collections = append(db.collections, &coll)
		fmt.Println(coll)
	}

	return &db, err
}

func (db *LocalDB) Logger() *slog.Logger {
	return db.log
}

func (db *LocalDB) GetCollections() []*Collection {
	return db.collections
}

func (db *LocalDB) Open(ctx context.Context, collection *Collection) (*os.File, error) {
	file, err := os.OpenFile(db.path+"/"+collection.name, os.O_RDWR, 0644)
	if err != nil {
		db.log.Error("Could not open file")
		file.Close()
	}
	collection.file = file
	return file, err
}

func (db *LocalDB) Write(ctx context.Context, collection *Collection, content []byte) (int, error) {
	if collection.file == nil {
		db.Open(ctx, collection)
	}
	bytes, err := collection.file.Write(content)
	if err != nil {
		db.log.Error("Could not write to file")
	}
	return bytes, err
}

func ReadFromCSV(ctx context.Context, path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	content, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	cByte := formatStrToJson(content)

	//jsonData, err := json.Marshal(cByte)
	//if err != nil {
	//	return nil, err
	//}

	return []byte(cByte), err
}

// Maybe implement template later
func formatStrToJson(content [][]string) string {
	fields := content[0]
	var cStr string = `[`

	var i int
	for i = 1; i < len(content); i++ {
		cStr +=
			`{"Id":` + `"` + uuid.New().String() + `"` +
				`,` + `"` + fields[0] + `"` + `:` + `"` + content[i][0] + `"` +
				`,` + `"` + fields[1] + `"` + `:` + content[i][1] +
				`,` + `"` + fields[2] + `"` + `:` + content[i][2] +
				`,` + `"` + fields[3] + `"` + `:` + content[i][3] + `},`
	}
	res, found := strings.CutSuffix(cStr, ",")
	if !found {
		panic(found)
	}
	res += `]`
	fmt.Println(res)

	return res
}
