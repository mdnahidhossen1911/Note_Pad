package userService

import (
	"note_pad/models"
	"note_pad/utils"
)

func (s *userService) Register(req *models.CreateUserRequest) (*models.RegisterResponce, error) {

	findEmail, err := s.repo.FindByEmail(req.Email)
	if findEmail != nil {
		return nil, models.ErrEmailExists
	}

	otp, err := utils.GenerateOTP(6)
	if err != nil {
		return nil, err
	}

	go func() {
		utils.SendOTPToEmail(otp, req.Email, req.Name, s.appPass, s.sendermail)

	}()

	u := &models.PandingUser{
		Name:     req.Name,
		Email:    req.Email,
		Password: utils.HashPassword(req.Password),
		Otp:      otp,
		IsOwner:  req.IsOwner,
	}

	return s.repo.CreatePanding(u)
}
