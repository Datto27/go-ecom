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
	g.POST("/", middlewares.JWTMiddleware(), h.AddProduct)
	g.GET("/:productId", h.GetPoroductById)
	g.GET("/", h.GetPoroducts)
	g.PATCH("/:productId", middlewares.JWTMiddleware(), h.UpdateProduct)
	g.PATCH("/update-image/:productId", middlewares.JWTMiddleware(), h.UpdateProductImage)
	g.DELETE("/:productId", middlewares.JWTMiddleware(), h.DeleteProduct)
}