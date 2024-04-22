package peer

import (
	"context"
	"log/slog"
	"os"

	"go.mongodb.org/mongo-driver/mongo/address"
)

//TODO: Create different types of LocalDB struct. (Document, user, sensor, more...)
//TODO: Use correct type of context in functions, not just TODO
//TODO: Create and use slog in functions
//TODO: More descriptive errors, maybe even just return them
//TODO: Maybe just use string instead of address package
//TODO: Add Defer functions to close files on crash
//TODO: Add keys to collections in write function

type LocalDBBase interface {
	ID() *Identity

	Open(ctx context.Context, collection *Collection) (*os.File, error)

	Create(ctx context.Context, name string) (Collection, error)

	Write(ctx context.Context, collection *Collection, content []byte) (int, error)

	GetAddress(ctx context.Context, name string) (address.Address, error)

	Logger() *slog.Logger
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
	path        address.Address
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
	file, err := os.Create(db.path.String() + "/" + name + ".json")
	if err != nil {
		db.log.Error("Could not create file")
	}
	file.Close()
	coll := Collection{name, nil, []byte(""), db.log}
	db.collections = append(db.collections, &coll)
	return &coll, err
}

func CreateLocalDatabase(ctx context.Context, name string, path address.Address, log slog.Logger) (*LocalDB, error) {
	err := os.Mkdir(path.String()+"/"+name, 0777)
	if err != nil {
		log.Error("Could not create directory")
	}
	var collections []*Collection
	db := LocalDB{name, path, collections, &log}
	return &db, err
}

func (db *LocalDB) Logger() *slog.Logger {
	return db.log
}

func (db *LocalDB) Open(ctx context.Context, collection *Collection) (*os.File, error) {
	file, err := os.Open(db.path.String() + "/" + collection.name)
	if err != nil {
		db.log.Error("Could not open file")
	}
	defer file.Close()
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
