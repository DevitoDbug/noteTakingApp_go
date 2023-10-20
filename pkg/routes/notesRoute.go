package routes

import (
	"github.com/DevitoDbug/noteTakingApp_go/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/note/", controllers.GetAllNotes).Methods("GET")
	r.HandleFunc("/note/{id}", controllers.GetNoteById).Methods("GET")
	r.HandleFunc("/note/", controllers.CreateNote).Methods("POST")
	r.HandleFunc("/note/{id}", controllers.UpdateNote).Methods("PUT")
	r.HandleFunc("/note/{id}", controllers.DeleteNote).Methods("Delete")
}
