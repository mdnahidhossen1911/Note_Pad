package usercontroller

import (
	"net/http"
	"note_pad/utils"

	"github.com/gin-gonic/gin"
)

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
