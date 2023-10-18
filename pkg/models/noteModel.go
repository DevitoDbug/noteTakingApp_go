package models

import (
	"github.com/DevitoDbug/noteTakingApp_go/pkg/config"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB

type Note struct {
	gorm.Model
	NoteId    string    `gorm:"" json:"note_id"`
	NoteTitle string    `json:"note_title"`
	NoteInfo  string    `json:"note_info"`
	NoteDate  time.Time `json:"note_date"`
}

func (n Note) CreateNote() *Note {
	db.NewRecord(n)
	db.Create(&n)
	return &n
}

func init() {
	config.CreateConnection()
}

func GetAllNotes() []Note {
	var notes []Note

	db.Find(&notes)
	return notes
}

func GetNoteById(ID int) (Note, *gorm.DB) {
	var note Note

	dbInstance := db.Where("NoteId=", ID)
	dbInstance.Find(&note)

	return note, dbInstance
}
