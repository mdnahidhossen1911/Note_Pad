package usercontroller

import (
	"net/http"
	"note_pad/utils"

	"github.com/gin-gonic/gin"
)

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
