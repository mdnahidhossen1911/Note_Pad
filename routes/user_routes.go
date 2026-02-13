package routes

import (
	"note_pad/controllers"
	"note_pad/middleware"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(rg *gin.RouterGroup, ctrl *controllers.UserController, jwtSecret string) {
	users := rg.Group("/users")

	// Public
	users.POST("", ctrl.Register)
	users.POST("/login", ctrl.Login)

	// Protected
	auth := users.Group("")
	auth.Use(middleware.AuthRequired(jwtSecret))
	{
		auth.GET("", ctrl.List)
		auth.GET("/:id", ctrl.GetByID)
		auth.PUT("/:id", ctrl.Update)
		auth.DELETE("/:id", ctrl.Delete)
	}
}
