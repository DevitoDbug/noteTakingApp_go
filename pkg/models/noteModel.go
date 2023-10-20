package models

import (
	"github.com/DevitoDbug/noteTakingApp_go/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Note struct {
	gorm.Model
	NoteTitle string `json:"note_title"`
	NoteInfo  string `json:"note_info"`
}

func (n Note) CreateNote() *Note {
	db.NewRecord(n)
	db.Create(&n)
	return &n
}

func init() {
	config.CreateConnection()
	db = config.GetConnection()
	db.AutoMigrate(&Note{})
}

func GetAllNotes() []Note {
	var notes []Note

	db.Find(&notes)
	return notes
}

func GetNoteById(ID int64) (*Note, *gorm.DB) {
	var note Note

	dbInstance := db.Where("ID=?", ID)
	dbInstance.Find(&note)

	return &note, dbInstance
}

func DeleteNoteById(ID int) Note {
	var note Note

	db.Where("ID=?", ID).Delete(&note)
	return note
}
