package userService

import (
	"note_pad/config"
	"note_pad/models"
	"note_pad/repositories"
	"note_pad/utils"
)

// UserService defines business logic for users.
type UserService interface {
	Register(req *models.CreateUserRequest) (*models.RegisterResponce, error)
	OtpVerification(req *models.OtpVerifyRequest) (string, error)
	RefreshToken(secret string) (string, error)
	Login(req *models.LoginRequest) (*models.LoginResponse, error)
	GetByID(id string) (*models.User, error)
	GetProfile(token string) (*models.User, error)
	List() ([]*models.User, error)
	Update(id string, u *models.User) (*models.User, error)
	Delete(id string) error
}

type userService struct {
	repo                 repositories.UserRepository
	jwtSecret            string
	jwtExpiryDays        int
	refreshjwtExpiryDays int
	appPass              string
	sendermail           string
}

func NewUserService(repo repositories.UserRepository, cfg *config.Config) UserService {
	return &userService{
		repo:                 repo,
		jwtSecret:            cfg.JwtSecureKey,
		jwtExpiryDays:        cfg.JwtExpiryDays,
		refreshjwtExpiryDays: cfg.RefreshJwtExpiryDays,
		appPass:              cfg.AppPass,
		sendermail:           cfg.SenderMail,
	}
}

func (s *userService) GetProfile(token string) (*models.User, error) {

	payload, err := utils.DecodeJWT(token, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	return s.repo.FindByID(payload.Sub)
}

func (s *userService) GetByID(id string) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) List() ([]*models.User, error) {
	return s.repo.List()
}

func (s *userService) Update(id string, u *models.User) (*models.User, error) {
	u.ID = id
	return s.repo.Update(u)
}

func (s *userService) Delete(id string) error {
	return s.repo.Delete(id)
}
