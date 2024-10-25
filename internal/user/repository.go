package user

import (
	"github.com/datto27/goecom/pkg/dtos"
	"github.com/datto27/goecom/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(data dtos.RegisterUserDto) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id string) (*models.User, error)
}

type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new Repository instance
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(data dtos.RegisterUserDto) error {
	user := &models.User{
		ID: uuid.New(),
		Email: data.Email,
		FirstName: data.FirstName,
		LastName: data.LastName,
		Password: data.Password,
	}
	result := r.db.Create(user);

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) GetUserByEmail(email string) (*models.User, error) {
	// var user models.User
	var user models.User

	// Use the Where clause to find the user by email
	result := r.db.Where("email = ?", email).Take(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return  &models.User{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (r *Repository) GetUserById(email string) (*models.User, error) {
	var user models.User
	return &user, nil
}
