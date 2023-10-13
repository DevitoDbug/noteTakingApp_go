package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	port := ":8080"

	r := mux.NewRouter()

	http.ListenAndServe(port, r)
}
