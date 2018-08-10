package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/wlwanpan/web-services/migration"
	"github.com/wlwanpan/web-services/user"
	mgo "gopkg.in/mgo.v2"
)

var (
	db *mgo.Session
)

func main() {

	port := ":" + os.Getenv("PORT")
	dbAddr := os.Getenv("DB_ADDR")

	router := mux.NewRouter()
	db, err := mgo.Dial(dbAddr)
	if err != nil {
		log.Fatal(err.Error())
	}

	user.SetRoutes(router)
	migration.MigrateData(db)

	server := &http.Server{
		Addr:        port,
		Handler:     router,
		ReadTimeout: 3 * time.Second,
	}

	log.Println("Server listening on port " + port)
	log.Fatal(server.ListenAndServe())
}
