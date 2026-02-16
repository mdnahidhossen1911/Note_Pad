package usercontroller

import (
	"net/http"
	"note_pad/models"
	"note_pad/utils"

	"github.com/gin-gonic/gin"
)

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
