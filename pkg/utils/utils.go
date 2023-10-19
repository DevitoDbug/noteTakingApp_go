package utils

import (
	"github.com/gorilla/mux"
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
