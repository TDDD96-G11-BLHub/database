package peer

import (
	"context"
	"encoding/json"
	"os"
	"time"
)

//TODO: implement support for adding structs for new sensordata structures.

type FilterFunc interface {
	filterMatch(doc DeepoidSenor) bool
}

type FilterTime struct {
	Value     time.Time
	Operation string
}

type FilterId struct {
	Value uint
}

type FilterFloat struct {
	Value     float64
	Operation string
	Field     string
}

type DeepoidSenor struct {
	id    uint
	Time  time.Time
	Roll  float64
	Pitch float64
	Yaw   float64
}

func LoadCollection(ctx context.Context, db *LocalDB, collection *Collection) ([]DeepoidSenor, error) {
	var res []DeepoidSenor
	doc, err := os.ReadFile(db.path.String() + "/" + collection.name)
	if err != nil {
		db.log.Error("Could not read from file")
	}
	er := json.Unmarshal([]byte(doc), &res)
	if er != nil {
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
	switch f.Operation {
	case "same":
		return doc.Time.Equal(f.Value)
	case "after":
		return doc.Time.After(f.Value)
	case "before":
		return doc.Time.Before(f.Value)
	default:
		return false
	}
}

func (f FilterId) filterMatch(doc DeepoidSenor) bool {
	return f.Value == doc.id
}

func (f FilterFloat) filterMatch(doc DeepoidSenor) bool {
	var val float64

	switch f.Field {
	case "Pitch":
		val = doc.Pitch
		break
	case "Yaw":
		val = doc.Yaw
		break
	case "Roll":
		val = doc.Roll
		break
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
	case "id":
		return FilterId{value.(uint)}
	case "Time":
		return FilterTime{value.(time.Time), operator}
	default:
		return nil
	}
}
