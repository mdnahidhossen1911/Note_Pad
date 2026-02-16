package routes

import (
	notecontroller "note_pad/controllers/note_controller"
	"note_pad/middleware"
	"note_pad/repositories"

	"github.com/gin-gonic/gin"
)

func registerNoteRoutes(rg *gin.RouterGroup, ctrl notecontroller.NoteController, userRepo repositories.UserRepository, jwtSecret string) {
	noteGr := rg.Group("/notes")

	note := noteGr.Group("")
	note.Use(middleware.AuthRequired(jwtSecret, userRepo))
	{
		note.POST("", ctrl.Create)
		note.GET("", ctrl.Get)
		note.GET("/profile", ctrl.GetById)
		note.PUT("/:id", ctrl.Update)
		note.DELETE("/:id", ctrl.Delete)
	}
}
