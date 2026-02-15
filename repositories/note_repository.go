package repositories

import (
	"note_pad/models"

	"gorm.io/gorm"
)

type noteRepository struct {
	db *gorm.DB
}

type NoteRepository interface {
	Create(note *models.NoteRequest) (*models.Note, error)
	List(Id string) ([]models.Note, error)
	Update(req *models.NoteUpdateRequest) (*models.Note, error)
	Delete(Id string) (string, error)
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return noteRepository{db: db}
}

// Create implements [NoteRepository].
func (n noteRepository) Create(note *models.NoteRequest) (*models.Note, error) {
	panic("unimplemented")
}

// Delete implements [NoteRepository].
func (n noteRepository) Delete(Id string) (string, error) {
	panic("unimplemented")
}

// List implements [NoteRepository].
func (n noteRepository) List(Id string) ([]models.Note, error) {
	panic("unimplemented")
}

// Update implements [NoteRepository].
func (n noteRepository) Update(req *models.NoteUpdateRequest) (*models.Note, error) {
	panic("unimplemented")
}
