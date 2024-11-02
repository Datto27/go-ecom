package dtos

import "github.com/google/uuid"

type CreateProductDto struct {
	UserID uuid.UUID `json:"userId" binding:"required"`
	Name string `json:"name" binding:"required"`
	Descritpion string `json:"description"`
	Quantity int `json:"quantity" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	Image string `json:"image"`
}

type ProductDto struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Image string `json:"image"`
}

type UpdateProductDto struct {
	ID uuid.UUID `json:"id" binding:"required"`
	UserID uuid.UUID `json:"userId" binding:"required"`
	Name string `json:"name" binding:"required"`
	Descritpion string `json:"description"`
	Quantity int `json:"quantity" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	Image string `json:"image"`
}