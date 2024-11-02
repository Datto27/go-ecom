package product

import (
	"fmt"

	"github.com/datto27/goecom/pkg/dtos"
	"github.com/datto27/goecom/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(data dtos.CreateProductDto) error
	FindProductById(id uuid.UUID) (*models.Product, error)
	FindProducts(skip int, limit int) ([]models.Product, error)
	UpdateProduct(userId uuid.UUID, data dtos.UpdateProductDto) (*models.Product, error)
	DeleteProduct(id uuid.UUID, userId uuid.UUID) (*models.Product, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) (*Repository) {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateProduct(data dtos.CreateProductDto) error {
	product := &models.Product{
		ID: uuid.New(),
		UserID: data.UserID,
		Name: data.Name,
		Description: data.Descritpion,
		Quantity: data.Quantity,
		Price: data.Price,
		Image: data.Image,
	}
	result := r.db.Create(product);

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) FindProductById(id uuid.UUID) (*models.Product, error) {
	var product models.Product

	result := r.db.Where("id = ?", id).Take(&product)

	fmt.Println("repo: ", result)

	if result.Error != nil {
		return nil, result.Error
	}

	return &models.Product{
		ID: product.ID,
		UserID: product.ID,
		Name: product.Name,
		Description: product.Description,
		Price: product.Price,
		Quantity: product.Quantity,
		Image: product.Image,
		UpdatedAt: product.UpdatedAt,
		CreatedAt: product.CreatedAt,
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
	result := r.db.Where("id = ? AND user_id = ?", data.ID, userId).Take(&product);

	if result.Error != nil {
		return nil, result.Error
	}

	result = r.db.Model(&product).Updates(models.Product{Name: data.Name, Description: data.Descritpion, Image: data.Image, Quantity: data.Quantity, Price: data.Price})

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (r *Repository) DeleteProduct(id uuid.UUID, userId uuid.UUID) (*models.Product, error) {
	var product models.Product
	result := r.db.Where("id = ? AND user_id = ?", id, userId).Delete(&product);

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}