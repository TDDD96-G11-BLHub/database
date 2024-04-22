package peer

import (
	"context"
	"encoding/json"
	"os"
	"time"
)

//TODO: implement support for adding structs for new sensordata structures.

type DeepoidSenor struct {
	id    int
	Time  time.Time
	Roll  float64
	Pitch float64
	Yaw   float64
}

func FindDocument(ctx context.Context, db *LocalDB, collection *Collection, id []byte) ([]byte, error) {
	var res []DeepoidSenor
	doc, err := os.ReadFile(db.path.String() + "/" + collection.name)
	if err != nil {
		db.log.Error("Could not read from file")
	}
	er := json.Unmarshal([]byte(doc), &res)
	if er != nil {
		db.log.Error("Could not Unmarshal document")
	}

}
