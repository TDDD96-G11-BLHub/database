package peerdb

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"log/slog"
	"os"
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
	ID() *Identity

	Open(ctx context.Context, collection *Collection) (*os.File, error)

	Create(ctx context.Context, name string) (Collection, error)

	Write(ctx context.Context, collection *Collection, content []byte) (int, error)

	GetAddress(ctx context.Context, name string) (string, error)

	Logger() *slog.Logger

	GetCollections() []*Collection
}

type Identity struct {
	iD         []byte
	pubKey     []byte
	signatures []*IDSignatures
}

type IDSignatures struct {
	iD     []byte
	pubKey []byte
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

func LoadLocalDatabase(ctx context.Context, name string, path string, log slog.Logger) *LocalDB {
	var collections []*Collection
	db := LocalDB{name: name, path: path + "/" + name, collections: collections, log: &log}
	return &db
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

	jsonData, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	return jsonData, err
}
