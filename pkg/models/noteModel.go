package models

import (
	"github.com/DevitoDbug/noteTakingApp_go/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Note struct {
	gorm.Model
	NoteId    string `gorm:"" json:"note_id"`
	NoteTitle string `json:"note_title"`
}

func (n Note) CreateNote() *Note {
	db.NewRecord(n)
	db.Create(&n)
	return &n
}

func init() {
	config.CreateConnection()

}
