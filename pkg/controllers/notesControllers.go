package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/DevitoDbug/noteTakingApp_go/pkg/models"
	"github.com/DevitoDbug/noteTakingApp_go/pkg/utils"
	"net/http"
)

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	notes := models.GetAllNotes()

	res, err := json.Marshal(notes)
	if err != nil {
		fmt.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		fmt.Print(err)
		return
	}
}

func GetNoteById(w http.ResponseWriter, r *http.Request) {
	ID, err := utils.GetIdParam(r)
	if err != nil {
		fmt.Print(err)
		return
	}

	note, dbInfo := models.GetNoteById(ID)
	if dbInfo.Error != nil {
		fmt.Print(dbInfo.Error)
		return
	}

	res, err := json.Marshal(note)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		fmt.Print(err)
		return
	}
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	newNote := &models.Note{}
	err := utils.ParseBody(r, newNote)
	if err != nil {
		fmt.Print(err)
		return
	}

	createdNote := newNote.CreateNote()

	res, err := json.Marshal(createdNote)
	if err != nil {
		fmt.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		fmt.Print(err)
	}
}
func UpdateNote(w http.ResponseWriter, r *http.Request) {

}
func DeleteNote(w http.ResponseWriter, r *http.Request) {

}
