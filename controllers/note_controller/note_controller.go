package notecontroller

import (
	noteservice "note_pad/services/note_service"

	"github.com/gin-gonic/gin"
)

type NoteController interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	GetById(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

func NewNoteController(srv noteservice.NoteService) NoteController {
	return notecontroller{
		service: srv,
	}
}

type notecontroller struct {
	service noteservice.NoteService
}

// Create implements [NoteController].
func (n notecontroller) Create(c *gin.Context) {
	panic("unimplemented")
}

// Delete implements [NoteController].
func (n notecontroller) Delete(c *gin.Context) {
	panic("unimplemented")
}

// Get implements [NoteController].
func (n notecontroller) Get(c *gin.Context) {
	panic("unimplemented")
}

// GetById implements [NoteController].
func (n notecontroller) GetById(c *gin.Context) {
	panic("unimplemented")
}

// Update implements [NoteController].
func (n notecontroller) Update(c *gin.Context) {
	panic("unimplemented")
}
