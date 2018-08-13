package user

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/wlwanpan/web-services/common"
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

// Check if filter value for device/os are valid.
func validateFilterQuery(s string, upperLimit int64) ([]int64, bool) {
	if s == "" {
		return []int64{}, true
	}

	filterArr := []int64{}
	splitStr := strings.Split(s, ",")
	for _, str := range splitStr {
		parsedInt := helper.StrToInt(str)
		if parsedInt < 0 || parsedInt > upperLimit {
			return []int64{}, false
		}
		filterArr = append(filterArr, parsedInt)
	}

	return filterArr, true
}

// Generate a CountQuery from request and call count.
func countQuery(db *mgo.Session, r *http.Request, uu bool) int {
	reqQuery := r.URL.Query()
	reqDevice := reqQuery.Get("device")
	reqOS := reqQuery.Get("os")

	filterDevice, valid := validateFilterQuery(reqDevice, 5)
	if !valid {
		return 0
	}
	filterOS, valid := validateFilterQuery(reqOS, 6)
	if !valid {
		return 0
	}

	q := &CountQuery{
		Os:     filterOS,
		Device: filterDevice,
		Unique: uu,
	}

	return q.Count(db)
}

func handlerHelper(option RespOptions, w http.ResponseWriter, r *http.Request) (int, error) {
	count := countQuery(option.Db, r, option.Uu)
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

// Handler for unique users request
func UniqueUserHandler(db *mgo.Session, w http.ResponseWriter, r *http.Request) (int, error) {
	option := RespOptions{
		Db: db,
		Uu: true,
	}
	return handlerHelper(option, w, r)
}

// Handler for loyal users request
func LoyalUserHandler(db *mgo.Session, w http.ResponseWriter, r *http.Request) (int, error) {
	option := RespOptions{
		Db: db,
		Uu: false,
	}
	return handlerHelper(option, w, r)
}
