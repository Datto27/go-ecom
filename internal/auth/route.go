package auth

import (
	"net/http"

	"github.com/datto27/goecom/internal/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository user.UserRepository
}

func NewHandler(respository user.UserRepository) * Handler {
	return &Handler{repository: respository}
}


func (h *Handler) Routes(g *gin.RouterGroup) {
	g.GET("/", h.test)
	g.POST("/register", h.HandleRegister)
	g.POST("/login", h.HandleLogin)
}

func (h *Handler) test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
