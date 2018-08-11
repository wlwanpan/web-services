package user

import (
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

func UseDB(h AppHandler, db *mgo.Session) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusCode, err := h(db, w, r)
		if err != nil {
			http.Error(w, err.Error(), statusCode)
		}
	}
}

func SetRoutes(r *mux.Router, db *mgo.Session) {
	r.HandleFunc("/unique-users", UseDB(UniqueUserHandler, db)).Methods("GET")
	r.HandleFunc("/loyal-users", UseDB(LoyalUserHandler, db)).Methods("GET")
}
