package user

import (
	"net/http"
	"strconv"

	utils "github.com/datto27/goecom/pkg"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Fetch Users
// @Summery Fetch Users
// @Description Get users list
// @Tags User
// @Produce json
// @Param skip query int true "Number of records to skip" default(0) "Get users with skip+limit queries"
// @Param limit query int true "Number of records to fetch" default(10) "Get users with skip+limit queries"
// @Success 200
// @Router /api/v1/users/ [get]
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

// Fetch User
// @Summery Fetch User
// @Description Get user details
// @Tags User
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} dtos.GetUserDto
// @Failure 404
// @Router /api/v1/users/{userId} [get]
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