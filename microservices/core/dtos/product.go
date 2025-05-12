package dtos

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type CreateProductDto struct {
	UserID string `form:"userId" binding:"required"`
	Name string `form:"name" binding:"required"`
	Descritpion string `form:"description"`
	Quantity int `form:"quantity" binding:"required"`
	Price float64 `form:"price" binding:"required"`
	Image *multipart.FileHeader `form:"image"`
}
type CreateProductDoc struct {
	UserID string `form:"userId" binding:"required"`
	Name string `form:"name" binding:"required"`
	Descritpion string `form:"description"`
	Quantity int `form:"quantity" binding:"required"`
	Price float64 `form:"price" binding:"required"`
	Image string `form:"image"`
}

type ProductDto struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Image *string `json:"image"`
}

type UpdateProductDto struct {
	ID uuid.UUID `json:"id" binding:"required"`
	UserID uuid.UUID `json:"userId" binding:"required"`
	Name string `json:"name" binding:"required"`
	Descritpion string `json:"description"`
	Quantity int `json:"quantity" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type UpdateProductImageDto struct {
	UserID string `form:"userId" binding:"required"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

type UpdateProductImageDoc struct {
	ProductID string `json:"productId" binding:"required" example:"12345"`
	Image     string `json:"image" example:"file.jpg"` // Represents the file name
}