package usercontroller

import (
	"net/http"
	"note_pad/utils"

	"github.com/gin-gonic/gin"
)

func (ctrl *userController) GetProfile(c *gin.Context) {

	token := utils.GetTokenFromHeader(c)

	user, err := ctrl.service.GetProfile(token)
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
