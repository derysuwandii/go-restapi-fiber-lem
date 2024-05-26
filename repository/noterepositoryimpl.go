package repository

import (
	"errors"
	"go-restapi-fiber-lem/data/request"
	"go-restapi-fiber-lem/helpers"
	"go-restapi-fiber-lem/models"
	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
	Db *gorm.DB
}

func (n *NoteRepositoryImpl) Save(note models.Note) {
	result := n.Db.Create(&note)
	helpers.ErrorPanic(result.Error)
}

func (n *NoteRepositoryImpl) Update(note models.Note) {
	var updateNote = request.UpdateNoteRequest{
		Id:      note.Id,
		Content: note.Content,
	}
	result := n.Db.Model(&note).Updates(updateNote)
	helpers.ErrorPanic(result.Error)
}

func (n *NoteRepositoryImpl) Delete(noteId int) {
	var note models.Note
	result := n.Db.Where("id=?", noteId).Delete(&note)
	helpers.ErrorPanic(result.Error)
}

func (n *NoteRepositoryImpl) FindById(noteId int) (models.Note, error) {
	var note models.Note
	result := n.Db.Find(&note, noteId)
	if result != nil {
		return note, nil
	} else {
		return note, errors.New("Note is not found")
	}
}

func (n *NoteRepositoryImpl) FindAll() []models.Note {
	var note []models.Note
	result := n.Db.Find(&note)
	helpers.ErrorPanic(result.Error)
	return note
}

func NewNoteRepositoryImpl(Db *gorm.DB) NoteRepository {
	return &NoteRepositoryImpl{Db: Db}
}
