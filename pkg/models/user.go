package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	Products []Product `gorm:"column:products;foreignKey:UserId;constraint:OnDelete:CASCADE" json:"products"`
	FirstName string `gorm:"column:firstName" json:"firstName"`
	LastName string	`gorm:"column:lastName" json:"lastName"`
	Email string	`gorm:"column:email;unique;not null" json:"email"`
	Password string	`gorm:"column:password" json:"password"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
}