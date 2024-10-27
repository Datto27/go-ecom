package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID uuid.UUID `gorm:"primary_key;column:id;default:uuid_generate_v4()"`
	UserId uuid.UUID `gorm:"column:UserId;not null" json:"userId"`
	User User `gorm:"references:UserId"`
	Name string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	Image string `gorm:"column:image" json:"image"`
	Price float64 `gorm:"column:price;not null" json:"price"`
	Quantity int `gorm:"column:quantity; not null" json:"quantity"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deletedAt" json:"deletedAt"`
}