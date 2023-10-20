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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, note); err != nil {
		return err
	}
	return nil
}
