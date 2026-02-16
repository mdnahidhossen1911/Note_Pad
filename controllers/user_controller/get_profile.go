package usercontroller

import (
	"net/http"
	"note_pad/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

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
