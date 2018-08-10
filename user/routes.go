package user

import (
	"github.com/gorilla/mux"
)

func SetRoutes(r *mux.Router) {

	r.HandleFunc("/unique-users", UserHandler).Methods("GET")
	r.HandleFunc("/loyal-users", UserHandler).Methods("GET")

}
