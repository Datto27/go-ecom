package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	FirstName string `gorm:"column:firstName" json:"firstName"`
	LastName string	`gorm:"column:lastName" json:"lastName"`
	Email string	`gorm:"column:email;unique;not null" json:"email"`
	Password string	`gorm:"column:password" json:"password"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
}