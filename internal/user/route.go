package user

import (
	"net/http"

	middlewares "github.com/datto27/goecom/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository UserRepository
}

func NewHandler(respository UserRepository) * Handler {
	return &Handler{repository: respository}
}

func (h *Handler) Routes(g *gin.RouterGroup) {
	g.Use(middlewares.JWTMiddleware())
	g.GET("/", h.test)
}

func (h *Handler) test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}