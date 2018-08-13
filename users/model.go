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

	// Query struct holds converted params queries
	CountQuery struct {
		Os     []int64
		Device []int64
		Unique bool
	}
)

// Insert new user doc to users collection in mongodb
func (u *User) Save(db *mgo.Session) error {
	return db.DB("").C(collection).Insert(u)
}

// Filter and count unique users
func uniqueCount(db *mgo.Session, q bson.M) int {
	var users []int
	err := db.DB("").C(collection).Find(q).Distinct("user", &users)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return len(users)
}

// Filter and count loyal (user visit > 9 within 1 month)
func loyalCount(db *mgo.Session, q bson.M) int {
	var result []interface{}

	pipeline := []bson.M{
		bson.M{"$match": q},
		bson.M{
			"$group": bson.M{
				"_id":      bson.M{"_id": "$user"},
				"datetime": bson.M{},
			},
		},
	}

	err := db.DB("").C(collection).Pipe(pipeline).All(&result)
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	return len(result)
}

// Query wrapper
func (q *CountQuery) Count(db *mgo.Session) int {
	bsonM := bson.M{}

	if len(q.Os) != 0 {
		bsonM["os"] = bson.M{"$in": q.Os}
	}
	if len(q.Device) != 0 {
		bsonM["device"] = bson.M{"$in": q.Device}
	}

	if q.Unique {
		return uniqueCount(db, bsonM)
	}
	return loyalCount(db, bsonM)
}
