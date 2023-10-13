package routes

import (
	"github.com/DevitoDbug/noteTakingApp_go/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/note/", controllers.GetAllNotes).Methods("GET")
	r.HandleFunc("/note/{id}", controllers.GetNotesById).Methods("GET")
	r.HandleFunc("/note/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/note/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/note/{id}", controllers.DeleteBook).Methods("Delete")
}
