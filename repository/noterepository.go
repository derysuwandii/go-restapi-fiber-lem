package repository

import "go-restapi-fiber-lem/models"

type NoteRepository interface {
	Save(note models.Note)
	Update(note models.Note)
	Delete(noteId int)
	FindById(noteId int) (models.Note, error)
	FindAll() []models.Note
}
