package controllers

import (
	"net/http"
	"note_pad/models"
	"note_pad/services"
	"note_pad/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// UserController handles all incoming HTTP requests for /users.
// Controller (C) — receives HTTP input, calls Service, returns HTTP response.
type UserController struct {
	service services.UserService
}

func NewUserController(svc services.UserService) *UserController {
	return &UserController{service: svc}
}

// Register godoc
// POST /users
func (ctrl *UserController) Register(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {

		if strings.Contains(err.Error(), "failed on the 'email' tag") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address."})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.service.Register(&req)
	if err != nil {

		switch err {
		case models.ErrEmailExists:
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
	}

	c.JSON(http.StatusCreated, utils.ApiResponse{
		Success: true,
		Message: "Account created. OTP has been sent to your email.",
		Data:    *user,
	})

}

func (ctrl *UserController) OtpVerification(c *gin.Context) {
	var req models.OtpVerifyRequest
	if error := c.ShouldBindJSON(&req); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error(),
		})
		return
	}

	token, err := ctrl.service.OtpVerification(&req)
	if err != nil {

		switch err {
		case models.ErrOTPInvalid:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		case models.ErrInvalidID:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		case models.ErrOTPExpired:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
	}
	c.JSON(http.StatusOK, models.LoginResponse{Token: token})

}

// Login godoc
// POST /users/login
func (ctrl *UserController) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.service.Login(&req)
	if err != nil {
		if err == models.ErrUserNotFound || err == models.ErrInvalidPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": true,
		"data":    *token,
	})
}

func (ctrl *UserController) RefrashToken(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		c.Abort()
		return
	}

	parts := strings.SplitN(header, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Format: Bearer <token>"})
		c.Abort()
		return
	}

	token, err := ctrl.service.RefreshToken(parts[1])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access-token": token,
	})
	// payload, err := utils.VerifyJWT(parts[1], jwtSecret , userRepo)
}

// GetByID godoc
// GET /users/:id
func (ctrl *UserController) GetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := ctrl.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// List godoc
// GET /users
func (ctrl *UserController) List(c *gin.Context) {
	users, err := ctrl.service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Update godoc
// PUT /users/:id
func (ctrl *UserController) Update(c *gin.Context) {
	id := c.Param("id")

	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := ctrl.service.Update(id, &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// Delete godoc
// DELETE /users/:id
func (ctrl *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
