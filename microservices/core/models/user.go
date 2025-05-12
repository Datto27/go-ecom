package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	Products []Product `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"products"`
	FirstName string `json:"firstName"`
	LastName string	`json:"lastName"`
	Email string	`gorm:"unique;not null" json:"email"`
	Password string	`json:"password"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}