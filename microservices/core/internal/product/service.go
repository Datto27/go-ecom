package product

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/datto27/goecom/microservices/core/dtos"
	"github.com/datto27/goecom/microservices/core/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// Add Product
// @Summery Add Product
// @Description Add new product
// @Tags Product
// @Accept multipart/form-data
// @Param product body dtos.CreateProductDoc true "Product create data"
// @Success 200
// @Failure 404
// @Router /api/v1/products [post]
func (h *Handler) AddProduct(c *gin.Context) {
	var productPayload dtos.CreateProductDto
	var newId = uuid.New()

	if err := c.ShouldBind(&productPayload); err != nil {
		utils.ApiError(c, http.StatusBadRequest, err.Error())
		return
	}

	file, err := c.FormFile("image")
	var imagePath string
	if err != nil || file == nil {
		productPayload.Image = nil
	} else {
		// save file to a local directory
		extension := strings.ToLower(filepath.Ext(file.Filename))
		filePath := fmt.Sprintf("uploads/%s%s", newId, extension)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			imagePath = filePath
		}
	}

	err = h.repository.CreateProduct(newId, productPayload, imagePath)

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully created"})
}

// Get Product
// @Summery Get Product
// @Description Get Product Details
// @Tags Product
// @Produce json
// @Param productId path string true "Product ID"
// @Success 200
// @Failure 404
// @Router /api/v1/products/{productId} [get]
func (h *Handler) GetPoroductById(c *gin.Context) {
	productId, err := uuid.Parse(c.Param("productId"))

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, "incorect id formation")
		return
	}

	product, err := h.repository.FindProductById(productId)

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

// Get Products
// @Summery Get Products
// @Description Get Products
// @Tags Product
// @Produce json
// @Param skip query int true "Number of records to skip" default(0)
// @Param limit query int true "Number of records to fetch" default(10)
// @Success 200
// @Failure 404
// @Router /api/v1/products [get]
func (h *Handler) GetPoroducts(c *gin.Context) {
	skip, err := strconv.Atoi(c.DefaultQuery("skip", "0"))
	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, "invalid skip value")
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, "invalid limit value")
		return
	}

	users, err := h.repository.FindProducts(skip, limit)

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": users})
}

// Update Product
// @Summery Update Product
// @Description Update product details
// @Tags Product
// @Produce json
// @Param productId path string true "Product ID"
// @Param product body dtos.UpdateProductDto true "Product update data"
// @Success 200
// @Failure 404
// @Router /api/v1/products/{productId} [patch]
func (h *Handler) UpdateProduct(c *gin.Context) {
	_, err := uuid.Parse(c.Param("productId"))
	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, "incorect id formation")
		return
	}

	var productPayload dtos.UpdateProductDto
	if err := c.ShouldBindJSON(&productPayload); err != nil {
		utils.ApiError(c, http.StatusBadRequest, err.Error())
		return
	}

	var userId uuid.UUID
	tokenStr := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	token, _ := utils.ParseJWT(tokenStr)

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if sub, ok := claims["sub"].(string); ok {
			userId, _ = uuid.Parse(sub)
		}
	}

	_, err = h.repository.UpdateProduct(userId, productPayload)

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product updated successfully"})
}

// Delete Product
// @Summery Delete Product
// @Description Delete product
// @Tags Product
// @Produce json
// @Param productId path string true "Product ID"
// @Success 200
// @Failure 404
// @Router /api/v1/products/{productId} [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {
	productId, err := uuid.Parse(c.Param("productId"))

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, "incorect id formation")
		return
	}

	var userId uuid.UUID
	tokenStr := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	token, _ := utils.ParseJWT(tokenStr)

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if sub, ok := claims["sub"].(string); ok {
			userId, _ = uuid.Parse(sub)
		}
	}

	_, err = h.repository.DeleteProduct(productId, userId)

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}

// Update Product Image
// @Summery Update Product Image
// @Description Update product details Image
// @Tags Product
// @Accept multipart/form-data
// @Param productId path string true "Product ID"
// @Param data body dtos.UpdateProductImageDoc true "Product image update with file"
// @Success 200
// @Failure 404
// @Router /api/v1/products/update-image/{productId} [patch]
func (h *Handler) UpdateProductImage(c *gin.Context) {
	var productPayload dtos.UpdateProductImageDto

	productId, err := uuid.Parse(c.Param("productId"))
	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, "incorect id formation")
		return
	}

	if err := c.ShouldBind(&productPayload); err != nil {
		utils.ApiError(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := uuid.Parse(productPayload.UserID)
	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, "incorect userId")
		return
	}

	var imagePath string
	file, err := c.FormFile("image")
	if err != nil || file == nil {
		productPayload.Image = nil
	} else {
		// save file to a local directory
		extension := strings.ToLower(filepath.Ext(file.Filename))
		filePath := fmt.Sprintf("uploads/%s%s", productId, extension)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			imagePath = filePath
		}
	}

	err = h.repository.UpdateProductImage(productId, userId, imagePath)

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, "could not update product image")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "image update successfully"})
		return
	}
}
