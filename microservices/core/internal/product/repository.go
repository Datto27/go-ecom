package product

import (
	"fmt"

	"github.com/datto27/goecom/microservices/core/dtos"
	"github.com/datto27/goecom/microservices/core/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(id uuid.UUID, data dtos.CreateProductDto, imagePath string) error
	FindProductById(id uuid.UUID) (*models.Product, error)
	FindProducts(skip int, limit int) ([]models.Product, error)
	UpdateProduct(userId uuid.UUID, data dtos.UpdateProductDto) (*models.Product, error)
	DeleteProduct(id uuid.UUID, userId uuid.UUID) (*models.Product, error)
	UpdateProductImage(id uuid.UUID, uid uuid.UUID, imagePath string) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateProduct(id uuid.UUID, data dtos.CreateProductDto, imagePath string) error {
	userId, err := uuid.Parse(data.UserID)
	if err != nil {
		return fmt.Errorf("user id parse error")
	}
	product := &models.Product{
		ID:          id,
		UserID:      userId,
		Name:        data.Name,
		Description: data.Descritpion,
		Quantity:    data.Quantity,
		Price:       data.Price,
		Image:       &imagePath,
	}
	result := r.db.Create(product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) FindProductById(id uuid.UUID) (*models.Product, error) {
	var product models.Product

	result := r.db.Where("id = ?", id).Take(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return &models.Product{
		ID:          product.ID,
		UserID:      product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Image:       product.Image,
		UpdatedAt:   product.UpdatedAt,
		CreatedAt:   product.CreatedAt,
	}, nil
}

func (r *Repository) FindProducts(skip int, limit int) ([]models.Product, error) {
	var products []models.Product

	result := r.db.Limit(limit).Offset(skip).Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (r *Repository) UpdateProduct(userId uuid.UUID, data dtos.UpdateProductDto) (*models.Product, error) {
	var product models.Product
	result := r.db.Where("id = ? AND user_id = ?", data.ID, userId).Take(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	result = r.db.Model(&product).Updates(models.Product{Name: data.Name, Description: data.Descritpion, Quantity: data.Quantity, Price: data.Price})

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (r *Repository) DeleteProduct(id uuid.UUID, userId uuid.UUID) (*models.Product, error) {
	var product models.Product
	result := r.db.Where("id = ? AND user_id = ?", id, userId).Delete(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (r *Repository) UpdateProductImage(id uuid.UUID, uid uuid.UUID, imagePath string) error {
	var product models.Product
	result := r.db.Where("id = ? AND user_id = ?", id, uid).Take(&product)

	if result.Error != nil {
		return result.Error
	}

	result = r.db.Model(&product).Updates(models.Product{Image: &imagePath})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
