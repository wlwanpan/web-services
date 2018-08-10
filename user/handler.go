package user

import (
	"encoding/json"
	"log"
	"net/http"
)

type Resp struct {
	Count uint `json:"count"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {

	resp := Resp{
		Count: 123,
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		log.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)

}
