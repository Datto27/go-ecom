package utils

import (
	"time"

	"github.com/datto27/goecom/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func ApiError(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{"message": msg})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var jwtSecret = []byte(config.ENVS.JWT_SECRET)

func GenerateJWT(id uuid.UUID) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id, 
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	})
	return claims.SignedString([]byte(jwtSecret))
}

func ParseJWT(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return jwtSecret, nil
	})

	return token, err
}
