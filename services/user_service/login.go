package userService

import (
	"note_pad/models"
	"note_pad/utils"
)

func (s *userService) Login(req *models.LoginRequest) (*models.TokenResponse, error) {
	u, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, models.ErrUserNotFound
	}

	if !utils.CheckPassword(req.Password, u.Password) {
		return nil, models.ErrInvalidPassword
	}

	token, _ := utils.GenerateJWT(u, utils.AccessToken, s.jwtSecret, s.jwtExpiryDays)
	refreshtoken, _ := utils.GenerateJWT(u, utils.RefreshToken, s.jwtSecret, s.refreshjwtExpiryDays)

	return &models.TokenResponse{
		Token:        token,
		RefreshToken: refreshtoken,
	}, nil

}
