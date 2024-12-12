package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID uuid.UUID `gorm:"primary_key;column:id;not null" json:"id"`
	UserID uuid.UUID `gorm:"not null" json:"userId"`
	Name string `gorm:"not null" json:"name"`
	Description string `json:"description"`
	Image *string `json:"image"`
	Price float64 `gorm:"not null" json:"price"`
	Quantity int `gorm:"not null" json:"quantity"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}