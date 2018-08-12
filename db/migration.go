package migration

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"time"

	"github.com/wlwanpan/web-services/common"
	"github.com/wlwanpan/web-services/users"
	mgo "gopkg.in/mgo.v2"
)

// Stores arr of errors for logging purposes
type Errors struct {
	errors []string
	count  uint
}

// Check and add error, increment error count
func (e *Errors) Add(err error) {
	if err != nil {
		e.errors = append(e.errors, err.Error())
		e.count++
	}
}

// Log accumulated error msg to console
func (e *Errors) Log() {
	if e.count != 0 {
		for _, err := range e.errors {
			log.Println(err)
		}
	}
	log.Println("Migration completed with %d errors", e.count)
}

// Migration function from .csv sample file to mongoDB
func MigrateData(db *mgo.Session) {
	errors := Errors{}
	defer errors.Log()
	dbSession := db.Copy()
	defer dbSession.Close()

	dataFile, err := os.Open("data.csv")
	errors.Add(err)
	r := csv.NewReader(bufio.NewReader(dataFile))

	log.Println("Starting data migration ...")
	for {
		entry, err := r.Read()
		if err == io.EOF {
			break
		} else {
			errors.Add(err)
		}

		newRecord := &user.User{
			Datetime: time.Unix(helper.StrToInt(entry[0]), 0),
			User:     helper.StrToUint(entry[1]),
			Os:       helper.StrToUint(entry[2]),
			Device:   helper.StrToUint(entry[3]),
		}

		log.Println("Inserting user: ", entry[1])
		errors.Add(newRecord.Save(dbSession))
	}
}
