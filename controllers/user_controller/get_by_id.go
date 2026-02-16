package usercontroller

import (
	"net/http"
	"note_pad/utils"

	"github.com/gin-gonic/gin"
)

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
