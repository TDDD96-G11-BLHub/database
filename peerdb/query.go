package peerdb

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

//TODO: implement support for adding structs for new sensordata structures.

type FilterFunc interface {
	filterMatch(doc DeepoidSenor) bool
}

type FilterTime struct {
	Value     string
	Operation string
}

type FilterId struct {
	Value uuid.UUID
}

type FilterFloat struct {
	Value     float64
	Operation string
	Field     string
}

type DeepoidSenor struct {
	Id    string
	Time  string
	Roll  float64
	Pitch float64
	Yaw   float64
}

func LoadCollection(ctx context.Context, db *LocalDB, collection *Collection) ([]DeepoidSenor, error) {
	var res []DeepoidSenor
	bytes, err := os.ReadFile(collection.file.Name())
	if err != nil {
		db.log.Error("Could not read from file")
		db.log.Error("Bytes read")
	}
	er := json.Unmarshal(bytes, &res)
	if er != nil {
		fmt.Println(er)
		db.log.Error("Could not Unmarshal document")
	}

	return res, err
}

func QueryCollection(ctx context.Context, db *LocalDB, docs []DeepoidSenor, f FilterFunc) []DeepoidSenor {
	var res []DeepoidSenor
	var i int

	for i = 0; i < len(docs); i++ {
		if f.filterMatch(docs[i]) {
			res = append(res, docs[i])
		}
	}
	if res == nil {
		db.log.Error("No matching documents")
	}
	return res
}

func (f FilterTime) filterMatch(doc DeepoidSenor) bool {
	filterTime, err := time.Parse(time.TimeOnly, f.Value)
	if err != nil {
		fmt.Println(err)
	}
	collTime, err := time.Parse(time.TimeOnly, doc.Time)
	if err != nil {
		fmt.Println(err)
	}
	switch f.Operation {
	case "same":
		return collTime.Equal(filterTime)
	case "after":
		return collTime.After(filterTime)
	case "before":
		return collTime.Before(filterTime)
	default:
		return false
	}
}

func (f FilterId) filterMatch(doc DeepoidSenor) bool {
	byteId, err := uuid.Parse(doc.Id)
	if err != nil {
		fmt.Println(err)
	}

	return f.Value == byteId
}

func (f FilterFloat) filterMatch(doc DeepoidSenor) bool {
	var val float64

	switch f.Field {
	case "Pitch":
		val = doc.Pitch
	case "Yaw":
		val = doc.Yaw
	case "Roll":
		val = doc.Roll
	default:
		return false
	}

	switch f.Operation {
	case "=":
		return val == f.Value
	case "<":
		return val < f.Value
	case ">":
		return val > f.Value
	case "<=":
		return val <= f.Value
	case ">=":
		return val >= f.Value
	default:
		return false
	}
}

func CreateFilter(field string, operator string, value any) FilterFunc {

	switch field {
	case "Yaw":
		return FilterFloat{value.(float64), operator, field}
	case "Pitch":
		return FilterFloat{value.(float64), operator, field}
	case "Roll":
		return FilterFloat{value.(float64), "=", field}
	case "Id":
		return FilterId{value.(uuid.UUID)}
	case "Time":
		return FilterTime{value.(string), operator}
	default:
		return nil
	}
}
