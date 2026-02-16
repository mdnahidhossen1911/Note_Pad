package routes

import (
	usercontroller "note_pad/controllers/user_controller"
	"note_pad/middleware"
	"note_pad/repositories"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(rg *gin.RouterGroup, ctrl usercontroller.UserController, userRepo repositories.UserRepository, jwtSecret string) {
	users := rg.Group("/users")

	// Public
	users.POST("", ctrl.Register)
	users.POST("/login", ctrl.Login)
	users.POST("/verification", ctrl.OtpVerification)
	users.GET("/refresh-token", ctrl.RefrashToken)

	// Protected
	auth := users.Group("")
	auth.Use(middleware.AuthRequired(jwtSecret, userRepo))
	{
		auth.GET("", ctrl.List)
		auth.GET("/profile", ctrl.GetProfile)
		auth.GET("/:id", ctrl.GetByID)
		auth.PUT("/:id", ctrl.Update)
		auth.DELETE("/:id", ctrl.Delete)
	}
}
