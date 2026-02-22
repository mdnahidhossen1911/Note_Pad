package noteservice

import (
	"note_pad/models"
	"note_pad/repositories"
	"note_pad/utils"
)

type NoteService interface {
	Create(note *models.NoteRequest, token string) (*models.Note, error)
	Get(token string) ([]*models.Note, error)
	GetProfile(token string) (*models.Note, error)
	Update(note models.NoteUpdateRequest) (*models.Note, error)
	Delete(id string) (string, error)
}

type noteService struct {
	repo      repositories.NoteRepository
	jwtSecret string
}

func NewNoteService(key string, repo repositories.NoteRepository) NoteService {
	return noteService{
		jwtSecret: key,
		repo:      repo,
	}
}

// Create implements [NoteService].
func (n noteService) Create(note *models.NoteRequest, token string) (*models.Note, error) {

	payload, err := utils.DecodeJWT(token, n.jwtSecret)

	if err != nil {
		return nil, err
	}

	noteData := &models.Note{
		UID:   payload.Sub,
		Title: note.Title,
		Body:  note.Body,
	}

	return n.repo.Create(noteData)

}

// Get implements [NoteService].
func (n noteService) Get(token string) ([]*models.Note, error) {

	payload, err := utils.DecodeJWT(token, n.jwtSecret)
	if err != nil {
		return nil, err
	}

	return n.repo.List(payload.Sub)
}

// Delete implements [NoteService].
func (n noteService) Delete(id string) (string, error) {

	return n.repo.Delete(id)

}

// GetProfile implements [NoteService].
func (n noteService) GetProfile(token string) (*models.Note, error) {
	panic("unimplemented")
}

// Update implements [NoteService].
func (n noteService) Update(note models.NoteUpdateRequest) (*models.Note, error) {
	return n.repo.Update(&note)
}
