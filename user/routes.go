package user

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

func Middleware(h AppHandler, db *mgo.Session) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer func() {
			log.Printf("[%s] %q %v", r.Method, r.URL.String(), time.Since(startTime))
		}()

		statusCode, err := h(db, w, r)
		if err != nil {
			http.Error(w, err.Error(), statusCode)
		}
	}
}

func SetRoutes(r *mux.Router, db *mgo.Session) {
	r.HandleFunc("/unique-users", Middleware(UniqueUserHandler, db)).Methods("GET")
	r.HandleFunc("/loyal-users", Middleware(LoyalUserHandler, db)).Methods("GET")
}
