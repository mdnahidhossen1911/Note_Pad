package usercontroller

import (
	"net/http"
	"note_pad/models"
	"note_pad/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

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
}
