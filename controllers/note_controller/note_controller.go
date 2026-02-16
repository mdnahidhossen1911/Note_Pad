package notecontroller

import (
	"net/http"
	"note_pad/models"
	noteservice "note_pad/services/note_service"
	"note_pad/utils"

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

func (n notecontroller) Create(c *gin.Context) {
	var req models.NoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	token := utils.GetTokenFromHeader(c)

	note, error := n.service.Create(&req, token)
	if error != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Message: error.Error(),
		})
	}

	c.JSON(http.StatusCreated, utils.ApiResponse{
		Success: true,
		Message: "Note Created",
		Data:    note,
	})

}

// Delete implements [NoteController].
func (n notecontroller) Delete(c *gin.Context) {
	panic("unimplemented")
}

// Get implements [NoteController].
func (n notecontroller) Get(c *gin.Context) {
	token := utils.GetTokenFromHeader(c)

	notes, err := n.service.Get(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Message: "Internal server error",
		})
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Data:    notes,
	})

}

// GetById implements [NoteController].
func (n notecontroller) GetById(c *gin.Context) {
	panic("unimplemented")
}

// Update implements [NoteController].
func (n notecontroller) Update(c *gin.Context) {
	panic("unimplemented")
}
