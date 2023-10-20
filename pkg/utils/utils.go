package utils

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetIdParam(r *http.Request) (int64, error) {
	param := mux.Vars(r)
	ID, err := strconv.ParseInt(param["ID"], 0, 0)
	if err != nil {
		return -1, err
	}
	return ID, nil
}

func ParseBody(r *http.Request, note interface{}) error {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), note); err != nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}
