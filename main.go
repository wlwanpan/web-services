package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/wlwanpan/web-services/common"
	"github.com/wlwanpan/web-services/db"
	"github.com/wlwanpan/web-services/users"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	serverAddr := helper.GetEnv("SERVER_ADDR", ":8080")
	dbAddr := helper.GetEnv("DB_ADDR", "localhost")
	migrate := helper.GetEnv("MIGRATE_ON_START", "false")

	dbDialInfo := &mgo.DialInfo{
		Addrs:    []string{dbAddr},
		Timeout:  60 * time.Second,
		Database: "service",
	}

	db, err := mgo.DialWithInfo(dbDialInfo)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.DB("").C("users").EnsureIndexKey("user")

	runMigration, err := strconv.ParseBool(migrate)
	if runMigration && err == nil {
		go migration.MigrateData(db)
	}

	router := mux.NewRouter()
	user.SetRoutes(router, db)

	server := &http.Server{
		Addr:        serverAddr,
		Handler:     router,
		ReadTimeout: 5 * time.Second,
	}

	log.Println("Server listening on " + serverAddr)
	log.Fatal(server.ListenAndServe())
}
