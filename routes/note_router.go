package routes

import (
	"note_pad/controllers"
	"note_pad/middleware"
	"note_pad/repositories"

	"github.com/gin-gonic/gin"
)

func registerNoteRoutes(rg *gin.RouterGroup, ctrl *controllers.UserController, userRepo repositories.UserRepository, jwtSecret string) {
	note := rg.Group("/note")

	auth := note.Group("")
	auth.Use(middleware.AuthRequired(jwtSecret, userRepo))
	{
		auth.GET("", ctrl.List)
		auth.POST("", ctrl.GetByID)
		auth.PUT("", ctrl.Update)
		auth.DELETE("", ctrl.Delete)
	}
}
