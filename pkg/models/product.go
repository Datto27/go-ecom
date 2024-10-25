package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id uuid.UUID `gorm:"primary_key;column:id;default:uuid_generate_v4()"`


	CreatedAt time.Time `gorm:"column:createdAt`
	UpdatedAt time.Time `gorm:"column:updatedAt`
}