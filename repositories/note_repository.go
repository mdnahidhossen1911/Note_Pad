package repositories

import (
	"fmt"
	"note_pad/models"

	"gorm.io/gorm"
)

type noteRepository struct {
	db *gorm.DB
}

type NoteRepository interface {
	Create(note *models.Note) (*models.Note, error)
	List(UId string) ([]*models.Note, error)
	Update(req *models.NoteUpdateRequest) (*models.Note, error)
	Delete(Id string) (string, error)
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return noteRepository{db: db}
}

// Create implements [NoteRepository].
func (n noteRepository) Create(note *models.Note) (*models.Note, error) {

	if err := n.db.Create(note).Error; err != nil {
		return nil, err
	}
	return note, nil

}

// List implements [NoteRepository].
func (n noteRepository) List(UId string) ([]*models.Note, error) {
	var notes []*models.Note

	if err := n.db.Where("uid = ?", UId).Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

// Delete implements [NoteRepository].
func (n noteRepository) Delete(id string) (string, error) {

	result := n.db.Delete(&models.Note{}, "id = ?", id)

	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		return "", fmt.Errorf("note not found")
	}

	return fmt.Sprintf("Note deleted successfully. ID: %s", id), nil
}

// Update implements [NoteRepository].
func (n noteRepository) Update(req *models.NoteUpdateRequest) (*models.Note, error) {

	var note models.Note
	if err := n.db.Model(note).Where("id=?", req.ID).Updates(map[string]interface{}{
		"body":  req.Body,
		"title": req.Title,
	}).Error; err != nil {
		return nil, err
	}

	if err := n.db.Where("id = ?", req.ID).Find(&note).Error; err != nil {
		return nil, err
	}
	return &note, nil

}
