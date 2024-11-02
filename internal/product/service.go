package product

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	utils "github.com/datto27/goecom/pkg"
	"github.com/datto27/goecom/pkg/dtos"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func (h *Handler) AddProduct(c *gin.Context) {
	var productPayload dtos.CreateProductDto

	if err := c.ShouldBindJSON(&productPayload); err != nil {
		fmt.Println("service: ", err)
		utils.ApiError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.repository.CreateProduct(productPayload)

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully created"})
}

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