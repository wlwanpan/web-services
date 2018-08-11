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
		Os     int64
		Device int64
		Unique bool
	}
)

func (u *User) Save(db *mgo.Session) error {
	return db.DB("").C(collection).Insert(u)
}

func (q *CountQuery) Count(db *mgo.Session) int {
	bsonM := bson.M{}
	if q.Os != -1 {
		bsonM["os"] = q.Os
	}
	if q.Device != -1 {
		bsonM["device"] = q.Device
	}

	count, err := db.DB("").C(collection).Find(bsonM).Count()
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return count
}
