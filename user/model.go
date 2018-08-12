package user

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collection string = "users"

type (
	// User mongo model < needs indexing
	User struct {
		Datetime time.Time `bson:"datetime"`
		User     uint64    `bson:"user"`
		Os       uint64    `bson:"os"`
		Device   uint64    `bson:"device"`
	}

	// Query
	CountQuery struct {
		Os     []int64
		Device []int64
		Unique bool
	}
)

func (u *User) Save(db *mgo.Session) error {
	return db.DB("").C(collection).Insert(u)
}

func (q *CountQuery) Count(db *mgo.Session) int {
	bsonM := bson.M{}

	if len(q.Os) != 0 {
		bsonM["os"] = bson.M{"$in": q.Os}
	}
	if len(q.Device) != 0 {
		bsonM["device"] = bson.M{"$in": q.Device}
	}

	if q.Unique == true {
		var u []int
		err := db.DB("").C(collection).Find(bsonM).Distinct("user", &u)
		if err != nil {
			log.Println(err.Error())
			return 0
		}
		return len(u)
	}

	return 0
}
