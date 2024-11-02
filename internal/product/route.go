package product

import (
	middlewares "github.com/datto27/goecom/middleware"
	"github.com/gin-gonic/gin"
)


type Handler struct {
	repository ProductRepository
}

func NewHandler(r ProductRepository) *Handler {
	return &Handler{
		repository: r,
	}
}

func (h *Handler) Routes(g *gin.RouterGroup) {
	g.Use(middlewares.JWTMiddleware())
	g.POST("/", h.AddProduct)
	g.GET("/:productId", h.GetPoroductById)
	g.GET("/", h.GetPoroducts)
	g.PATCH("/:productId", h.UpdateProduct)
	g.DELETE("/:productId", h.DeleteProduct)
}