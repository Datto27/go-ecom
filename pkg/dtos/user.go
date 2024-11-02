package dtos

import (
	"time"

	"github.com/google/uuid"
)

type GetUserDto struct {
	ID uuid.UUID `json:"id"`
	Products []ProductDto `json:"products"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"emial"`
	Password string `json:"passowrd"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}