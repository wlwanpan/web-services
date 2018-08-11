package migration

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"time"

	"github.com/wlwanpan/web-services/parser"
	"github.com/wlwanpan/web-services/user"
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

	filePath := os.Getenv("CSV_PATH")
	if filePath == "" {
		filePath = "data.csv"
	}

	dataFile, err := os.Open(filePath)
	errors.Add(err)
	r := csv.NewReader(bufio.NewReader(dataFile))

	log.Println("Starting data migration from: ", filePath)
	for {
		entry, err := r.Read()
		if err == io.EOF {
			break
		} else {
			errors.Add(err)
		}

		newRecord := &user.User{
			Datetime: time.Unix(parser.StrToInt(entry[0]), 0),
			User:     parser.StrToUint(entry[1]),
			Os:       parser.StrToUint(entry[2]),
			Device:   parser.StrToUint(entry[3]),
		}

		log.Println("Inserting User: ", entry[1])
		errors.Add(newRecord.Save(dbSession))
	}
}
