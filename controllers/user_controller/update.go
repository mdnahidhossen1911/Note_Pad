package usercontroller

import (
	"net/http"
	"note_pad/models"
	"note_pad/utils"

	"github.com/gin-gonic/gin"
)

func (ctrl *userController) Update(c *gin.Context) {
	id := c.Param("id")

	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: "Invalid payload",
		})
		return
	}

	if u.Name == "" {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: "Name is requird. Key:name",
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
