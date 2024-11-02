package user

import (
	"github.com/datto27/goecom/pkg/dtos"
	"github.com/datto27/goecom/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(data dtos.RegisterUserDto) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindUserById(id uuid.UUID) (*dtos.GetUserDto, error)
	FindUsers(skip int, limit int) ([]models.User, error)
}

type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new Repository instance
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(data dtos.RegisterUserDto) (*models.User, error){
	user := &models.User{
		ID: uuid.New(),
		Email: data.Email,
		FirstName: data.FirstName,
		LastName: data.LastName,
		Password: data.Password,
	}
	result := r.db.Create(user);

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *Repository) FindUserByEmail(email string) (*models.User, error) {
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

func (r *Repository) FindUserById(id uuid.UUID) (*dtos.GetUserDto, error) {
	var user models.User

	// Use the Where clause to find the user by email
	// result := r.db.Where("id = ?", id).Take(&user)

	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	err := r.db.Preload("Products", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Name", "Price", "Image", "UserID")
		}).First(&user, "id = ?", id).Error

	if err != nil {
    return nil, err
	}

	products := make([]dtos.ProductDto, len(user.Products))
	for i, product := range user.Products {
			products[i] = dtos.ProductDto{
					ID: product.ID,
					Name: product.Name,
					Price: product.Price,
					Image: product.Image,
			}
	}

	return  &dtos.GetUserDto{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Products:  products,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (r *Repository) FindUsers(skip int, limit int) ([]models.User, error) {
	var users []models.User

	result := r.db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}