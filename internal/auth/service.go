package auth

import (
	"net/http"
	"time"

	utils "github.com/datto27/goecom/pkg"
	"github.com/datto27/goecom/pkg/dtos"
	"github.com/gin-gonic/gin"
)


func (h *Handler) HandleRegister(c *gin.Context) {
	var userPayload dtos.RegisterUserDto

	if err := c.ShouldBindJSON(&userPayload); err != nil {
		utils.ApiError(c, http.StatusBadRequest, err.Error())
		return
	}

	// check if user exists before create new one
	user, err := h.repository.FindUserByEmail(userPayload.Email)
	if (user != nil) {
		utils.ApiError(c, http.StatusBadRequest, "user already with same email")
		return
	}

	pass, err := utils.HashPassword(userPayload.Password)

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, "password hashing problem")
		return
	}

	userPayload.Password = pass
	
	user, err = h.repository.CreateUser(userPayload)

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, "user creation problem")
		return
	}

	token, err := utils.GenerateJWT(user.ID)

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, "token generation error")
		return
	}

	c.SetCookie(
		"jwt",
		token,
		int(time.Now().Add(time.Hour * 24).Unix()),
		"/",
		"",
		true,
		true,
	)


	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
 
func (h *Handler) HandleLogin(c *gin.Context) {
	var loginPayload dtos.LoginUserDto

	if err := c.ShouldBindJSON(&loginPayload); err != nil {
		utils.ApiError(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.repository.FindUserByEmail(loginPayload.Email)

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, "user not found")
		return
	}

	isPassValid := utils.VerifyPassword(loginPayload.Password, user.Password)

	if !isPassValid {
		utils.ApiError(c, http.StatusBadRequest, "password is incorrect")
		return
	}

	token, err := utils.GenerateJWT(user.ID)

	if err != nil {
		utils.ApiError(c, http.StatusBadRequest, "token generation error")
		return
	}

	c.SetCookie(
		"jwt",
		token,
		int(time.Now().Add(time.Hour * 24 * 7).Unix()),
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"message": "success", "user": user})
}
