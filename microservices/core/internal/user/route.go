package user

import (
	middlewares "github.com/datto27/goecom/microservices/core/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository UserRepository
}

func NewHandler(respository UserRepository) *Handler {
	return &Handler{repository: respository}
}

func (h *Handler) Routes(g *gin.RouterGroup) {
	g.GET("/", h.GetUsers)
	g.GET("/:userId", middlewares.JWTMiddleware(), h.GetUserById)
}
