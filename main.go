package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/wlwanpan/web-services/migration"
	"github.com/wlwanpan/web-services/user"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	port := os.Getenv("SERVER_ADDR")
	if port == "" {
		port = ":8080"
	}
	dbAddr := os.Getenv("DB_ADDR")
	if dbAddr == "" {
		dbAddr = "localhost"
	}
	runMigration := *flag.Bool("m", false, "migrate data before running server")

	dbDialInfo := &mgo.DialInfo{
		Addrs:    []string{dbAddr},
		Timeout:  60 * time.Second,
		Database: "service",
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

	log.Println("Server listening on " + port)
	log.Fatal(server.ListenAndServe())
}
