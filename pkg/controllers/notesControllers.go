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
func CreateBook(w http.ResponseWriter, r *http.Request) {

}
func UpdateBook(w http.ResponseWriter, r *http.Request) {

}
func DeleteBook(w http.ResponseWriter, r *http.Request) {

}
