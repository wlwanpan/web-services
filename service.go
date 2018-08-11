package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/wlwanpan/web-services/migration"
	"github.com/wlwanpan/web-services/user"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	port := ":" + *flag.String("p", "3001", "listening port")
	dbAddr := *flag.String("db", "localhost", "mongodb addr")
	runMigration := *flag.Bool("m", false, "migrate data before running server")

	dbDialInfo := &mgo.DialInfo{
		Addrs:    []string{dbAddr},
		Database: "users", // to change to service
	}

	db, err := mgo.DialWithInfo(dbDialInfo)
	if err != nil {
		log.Fatal(err.Error())
	}
	if runMigration == true {
		migration.MigrateData(db)
	}

	router := mux.NewRouter()
	user.SetRoutes(router, db)

	server := &http.Server{
		Addr:        port,
		Handler:     router,
		ReadTimeout: 3 * time.Second,
	}

	log.Println("Server listening on port " + port)
	log.Fatal(server.ListenAndServe())
}
