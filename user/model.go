package user

import (
	"time"
)

type User struct {
	Datetime time.Time `json:"datetime" bson:"datetime"`
	User     uint64    `json:"user" bson:"user"`
	Os       uint64    `json:"os" bson:"os"`
	Device   uint64    `json:"device" bson:"device"`
}
