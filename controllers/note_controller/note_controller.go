package notecontroller

import (
	"net/http"
	"note_pad/models"
	noteservice "note_pad/services/note_service"
	"note_pad/utils"
	"strings"

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
	id := c.Param("id")

	msg, err := n.service.Delete(id)

	if err != nil {

		if strings.Contains(err.Error(), "note not found") {
			c.JSON(http.StatusNotFound, utils.ApiResponse{
				Success: false,
				Message: "Note is not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: msg,
	})

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

	id := c.Param("id")

	var note models.NoteUpdateRequest

	if c.ShouldBindJSON(&note) != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: "Invalid payload",
		})
		return
	}

	if note.Title == "" {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: "Requird. key:title",
		})
		return
	}

	if note.Body == "" {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: "Requird. key:body",
		})
		return
	}

	note.ID = id

	u_note, err := n.service.Update(note)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Data:    u_note,
	})

}
