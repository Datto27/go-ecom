package user

import (
	"net/http"
	"strconv"

	utils "github.com/datto27/goecom/pkg"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetUsers(c *gin.Context) {
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
	
	users, err := h.repository.FindUsers(skip, limit)

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *Handler) GetUserById(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("userId"))

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, "invalid user id")
		return
	}
	user, err := h.repository.FindUserById(userId)

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}