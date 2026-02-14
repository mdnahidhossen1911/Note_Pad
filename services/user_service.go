package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"note_pad/config"
	"note_pad/models"
	"note_pad/repositories"
	"note_pad/utils"
	"strings"
	"time"
)

// UserService defines business logic for users.
type UserService interface {
	Register(req *models.CreateUserRequest) (*models.RegisterResponce, error)
	OtpVerification(req *models.OtpVerifyRequest) (string, error)
	RefreshToken(secret string) (string, error)
	Login(req *models.LoginRequest) (*models.LoginResponse, error)
	GetByID(id string) (*models.User, error)
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

// RefreshToken implements [UserService].
func (s *userService) RefreshToken(token string) (string, error) {

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid token")
	}

	msg := parts[0] + "." + parts[1]
	mac := hmac.New(sha256.New, []byte(s.jwtSecret))
	mac.Write([]byte(msg))
	expected := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	if expected != parts[2] {
		return "", fmt.Errorf("invalid signature")
	}

	raw, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("invalid payload")
	}

	var p utils.JWTPayload
	if err := json.Unmarshal(raw, &p); err != nil {
		return "", fmt.Errorf("malformed payload")
	}

	u, err := s.repo.FindByID(p.Sub)

	if err != nil {
		return "", fmt.Errorf("User does not exist")
	}

	if p.Type != utils.RefreshToken {
		return "", fmt.Errorf("this endpoint requires an refresh token")
	}

	// Check token expiration
	if p.Exp > 0 && time.Now().Unix() > p.Exp {
		return "", fmt.Errorf("token expired")
	}

	return utils.GenerateJWT(u, utils.AccessToken, s.jwtSecret, s.jwtExpiryDays)

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

// OtpVerification implements [UserService].
func (s *userService) OtpVerification(req *models.OtpVerifyRequest) (string, error) {
	if len(req.Otp) != 6 {
		return "", models.ErrOTPInvalid
	}

	tuser, err := s.repo.PandingUserFindById(req.Uid)

	if err != nil {
		return "", models.ErrInvalidID
	}

	if req.Otp != tuser.Otp {
		return "", models.ErrOTPInvalid
	}

	u := &models.User{
		Name:     tuser.Name,
		Email:    tuser.Email,
		Password: tuser.Password,
		IsOwner:  tuser.IsOwner,
	}

	isValid := time.Since(tuser.CreatedAt).Seconds() <= 120
	if !isValid {
		return "", models.ErrOTPExpired
	}

	user, err := s.repo.Create(u)
	if err != nil {
		return "", err
	}

	err = s.repo.DeletePandingUser(req.Uid)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(user, utils.AccessToken, s.jwtSecret, s.jwtExpiryDays)
	return token, err

}

func (s *userService) Login(req *models.LoginRequest) (*models.LoginResponse, error) {
	u, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, models.ErrUserNotFound
	}

if !utils.CheckPassword(req.Password, u.Password) {
    return nil, models.ErrInvalidPassword
}

	token, _ := utils.GenerateJWT(u, utils.AccessToken, s.jwtSecret, s.jwtExpiryDays)
	refreshtoken, _ := utils.GenerateJWT(u, utils.RefreshToken, s.jwtSecret, s.refreshjwtExpiryDays)

	return &models.LoginResponse{
		Token:        token,
		RefreshToken: refreshtoken,
	}, nil

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
