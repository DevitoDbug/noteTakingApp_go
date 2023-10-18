package main

import (
	"fmt"
	"github.com/DevitoDbug/noteTakingApp_go/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	port := ":8080"

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)

	log.Printf("Starting sever at port %v", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		fmt.Print(err)
	}
}
