package usercontroller

import (
	userService "note_pad/services/user_service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(c *gin.Context)
	OtpVerification(c *gin.Context)
	Login(c *gin.Context)
	RefrashToken(c *gin.Context)

	GetByID(c *gin.Context)
	GetProfile(c *gin.Context)
	List(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type userController struct {
	service userService.UserService
}

func NewUserController(svc userService.UserService) UserController {
	return &userController{
		service: svc,
	}
}
