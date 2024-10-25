package dtos

type RegisterUserDto struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName string	`json:"lastName" binding:"required"`
	Email string	`json:"email" binding:"required"`
	Password string	`json:"password" binding:"required"`
}

type LoginUserDto struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}