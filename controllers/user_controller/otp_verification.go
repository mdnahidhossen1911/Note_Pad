package usercontroller

import (
	"net/http"
	"note_pad/models"
	"note_pad/utils"

	"github.com/gin-gonic/gin"
)

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
