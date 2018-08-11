package user

import (
	"encoding/json"
	"net/http"

	"github.com/wlwanpan/web-services/parser"
	mgo "gopkg.in/mgo.v2"
)

type (
	// Response struct
	Resp struct {
		Count int `json:"count"`
	}
	// Response option
	RespOptions struct {
		Db *mgo.Session
		Uu bool
	}
)

type AppHandler func(db *mgo.Session, w http.ResponseWriter, r *http.Request) (int, error)

func validateFilterQuery(s string) (int64, bool) {
	if s == "" {
		return -1, true
	}

	parsedInt := parser.StrToInt(s)
	if parsedInt < 0 || parsedInt > 5 {
		return 0, false
	}
	return parsedInt, true
}

func genQuery(db *mgo.Session, r *http.Request, uu bool) int {
	reqQuery := r.URL.Query()
	reqDevice := reqQuery.Get("device")
	reqOS := reqQuery.Get("os")

	intDevice, valid := validateFilterQuery(reqDevice)
	if valid == false {
		return 0
	}
	intOS, valid := validateFilterQuery(reqOS)
	if valid == false {
		return 0
	}

	q := &CountQuery{
		Os:     intOS,
		Device: intDevice,
		Unique: uu,
	}

	return q.Count(db)
}

func handlerHelper(option RespOptions, w http.ResponseWriter, r *http.Request) (int, error) {
	count := genQuery(option.Db, r, option.Uu)
	resp := Resp{Count: count}
	respJson, err := json.Marshal(resp)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
	return http.StatusOK, nil
}

func UniqueUserHandler(db *mgo.Session, w http.ResponseWriter, r *http.Request) (int, error) {
	option := RespOptions{
		Db: db,
		Uu: true,
	}
	return handlerHelper(option, w, r)
}

func LoyalUserHandler(db *mgo.Session, w http.ResponseWriter, r *http.Request) (int, error) {
	option := RespOptions{
		Db: db,
		Uu: false,
	}
	return handlerHelper(option, w, r)
}
