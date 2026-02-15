package controllers

import (
	"net/http"
	"note_pad/models"
	userService "note_pad/services/user_service"
	"note_pad/utils"
	"strings"

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

func (ctrl *userController) Register(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {

		if strings.Contains(err.Error(), "failed on the 'email' tag") {
			c.JSON(http.StatusBadRequest, utils.ApiResponse{
				Success: false,
				Message: "Invalid email address.",
			})
			return
		}

		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	user, err := ctrl.service.Register(&req)
	if err != nil {

		switch err {
		case models.ErrEmailExists:
			c.JSON(http.StatusConflict, utils.ApiResponse{
				Success: false,
				Message: err.Error(),
			})
			return

		default:
			c.JSON(http.StatusInternalServerError, utils.ApiResponse{
				Success: false,
				Message: "Internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusCreated, utils.ApiResponse{
		Success: true,
		Message: "Account created. OTP has been sent to your email.",
		Data:    *user,
	})

}

func (ctrl *userController) OtpVerification(c *gin.Context) {
	var req models.OtpVerifyRequest
	if error := c.ShouldBindJSON(&req); error != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: error.Error(),
		})
		return
	}

	token, err := ctrl.service.OtpVerification(&req)
	if err != nil {

		switch err {
		case models.ErrOTPInvalid:
			c.JSON(http.StatusBadRequest,
				utils.ApiResponse{
					Success: false,
					Message: err.Error(),
				})
			return

		case models.ErrInvalidID:
			c.JSON(http.StatusBadRequest,
				utils.ApiResponse{
					Success: false,
					Message: err.Error(),
				})
			return

		case models.ErrOTPExpired:
			c.JSON(http.StatusBadRequest,
				utils.ApiResponse{
					Success: false,
					Message: err.Error(),
				})
			return

		default:
			c.JSON(http.StatusBadRequest,
				utils.ApiResponse{
					Success: false,
					Message: "Internal server error",
				})
			return
		}
	}
	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Otp Verification Successful",
		Data:    models.TokenResponse{Token: token},
	})

}

func (ctrl *userController) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	token, err := ctrl.service.Login(&req)
	if err != nil {
		if err == models.ErrUserNotFound || err == models.ErrInvalidPassword {
			c.JSON(http.StatusUnauthorized, utils.ApiResponse{
				Success: false,
				Message: "Invalid credentials",
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
		Message: "Login Successful",
		Data:    *token,
	})
}

func (ctrl *userController) RefrashToken(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.JSON(http.StatusUnauthorized, utils.ApiResponse{
			Success: false,
			Message: "Authorization header required",
		})
		c.Abort()
		return
	}

	parts := strings.SplitN(header, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
		c.JSON(http.StatusUnauthorized, utils.ApiResponse{
			Success: false,
			Message: "Format: Bearer <token>",
		})
		c.Abort()
		return
	}

	token, err := ctrl.service.RefreshToken(parts[1])
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Access Token Genaret Successful",
		Data:    models.TokenResponse{Token: token},
	})
	// payload, err := utils.VerifyJWT(parts[1], jwtSecret , userRepo)
}

func (ctrl *userController) GetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := ctrl.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ApiResponse{
			Success: false,
			Message: "User not found",
		})
		return
	}
	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Data:    user,
	})
}

func (ctrl *userController) GetProfile(c *gin.Context) {

	header := c.GetHeader("Authorization")
	if header == "" {
		c.JSON(http.StatusUnauthorized, utils.ApiResponse{
			Success: false,
			Message: "Authorization header required",
		})
		c.Abort()
		return
	}

	parts := strings.SplitN(header, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
		c.JSON(http.StatusUnauthorized, utils.ApiResponse{
			Success: false,
			Message: "Format: Bearer <token>",
		})
		c.Abort()
		return
	}

	user, err := ctrl.service.GetProfile(parts[1])
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Data:    "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Data:    user,
	})

}

func (ctrl *userController) List(c *gin.Context) {
	users, err := ctrl.service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Data:    users,
	})
}

func (ctrl *userController) Update(c *gin.Context) {
	id := c.Param("id")

	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	updated, err := ctrl.service.Update(id, &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (ctrl *userController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "User deleted successfully",
	})
}
