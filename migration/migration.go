package migration

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/wlwanpan/web-services/user"
	mgo "gopkg.in/mgo.v2"
)

// Converts a string to uint64.
// If error in parsing, returns 0 and log the error
func parseStrToUint(s string) uint64 {
	parsedS, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Println("Error Parsing to Uint: ", s)
		return 0
	}
	return parsedS
}

// Migration function from .csv sample file to mongoDB
func MigrateData(db *mgo.Session) error {
	dbSession := db.Copy()
	defer dbSession.Close()

	filePath := os.Getenv("CSV_PATH")
	if filePath == "" {
		filePath = "data.csv"
	}

	log.Println("Starting data migration from: ", filePath)

	dataFile, err := os.Open(filePath)
	if err != nil {
		return err
	}

	r := csv.NewReader(bufio.NewReader(dataFile))

	for {
		entry, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		parsedTime, err := time.Parse(time.UnixDate, entry[0])
		if err != nil {
			log.Println("Error Parsing time obj of User: ", entry[1])
			parsedTime = time.Now()
		}

		record := &user.User{
			Datetime: parsedTime,
			User:     parseStrToUint(entry[1]),
			Os:       parseStrToUint(entry[2]),
			Device:   parseStrToUint(entry[3]),
		}

		dbSession.DB("services").C("users").Insert(record)
	}

	return nil
}
