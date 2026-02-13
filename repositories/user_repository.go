package repositories

import (
	"note_pad/models"

	"gorm.io/gorm"
)

// UserRepository defines the DB contract for users.
type UserRepository interface {
	Create(u *models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByID(id string) (*models.User, error)
	List() ([]*models.User, error)
	Update(u *models.User) (*models.User, error)
	Delete(id string) error

	CreatePanding(u *models.PandingUser) (*models.RegisterResponce, error)
	PandingUserFindById(id string) (*models.PandingUser, error)
	DeletePandingUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

// CreatePanding implements [UserRepository].
func (r *userRepository) CreatePanding(u *models.PandingUser) (*models.RegisterResponce, error) {
	if error := r.db.Create(u).Error; error != nil {
		return nil, error
	}
	return &models.RegisterResponce{
		UID: u.ID,
	}, nil

}

// DeletePandingUser implements [UserRepository].
func (r *userRepository) DeletePandingUser(id string) error {
	result := r.db.Delete(&models.PandingUser{}, "id = ?", id)
	return result.Error
}

// PandingUserFindById implements [UserRepository].
func (r *userRepository) PandingUserFindById(id string) (*models.PandingUser, error) {
	var u models.PandingUser
	if err := r.db.First(&u, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(u *models.User) (*models.User, error) {
	if err := r.db.Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var u models.User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) FindByID(id string) (*models.User, error) {
	var u models.User
	if err := r.db.First(&u, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) List() ([]*models.User, error) {
	var users []*models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Update(u *models.User) (*models.User, error) {
	if err := r.db.Model(u).Updates(map[string]interface{}{
		"name":  u.Name,
		"email": u.Email,
	}).Error; err != nil {
		return nil, err
	}
	// Reload to get updated timestamps
	if err := r.db.First(u, "id = ?", u.ID).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (r *userRepository) Delete(id string) error {
	result := r.db.Delete(&models.User{}, "id = ?", id)
	return result.Error
}
