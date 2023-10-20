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
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetNoteById(w http.ResponseWriter, r *http.Request) {
	ID, err := utils.GetIdParam(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	note, dbInfo := models.GetNoteById(ID)
	if dbInfo.Error != nil {
		fmt.Println(dbInfo.Error)
		return
	}

	res, err := json.Marshal(note)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	newNote := &models.Note{}
	err := utils.ParseBody(r, newNote)
	if err != nil {
		fmt.Println(err)
		return
	}

	createdNote := newNote.CreateNote()

	res, err := json.Marshal(createdNote)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		fmt.Println(err)
	}
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	var newNote = &models.Note{}
	utils.ParseBody(r, newNote)
	ID, err := utils.GetIdParam(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	note, dbInfo := models.GetNoteById(ID)
	if dbInfo.Error != nil {
		fmt.Println(dbInfo.Error)
		return
	}

	//Checking for changes
	if newNote.NoteInfo != "" {
		note.NoteInfo = newNote.NoteInfo
	}
	if newNote.NoteTitle != "" {
		note.NoteTitle = newNote.NoteTitle
	}

	//saving changes
	dbInfo.Save(&note)

	res, err := json.Marshal(note)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		fmt.Println(err)
	}
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	ID, err := utils.GetIdParam(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	deletedNote := models.DeleteNoteById(ID)

	res, err := json.Marshal(deletedNote)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(res); err == nil {
		fmt.Println(err)
		return
	}

}
