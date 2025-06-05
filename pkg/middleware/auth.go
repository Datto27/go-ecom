package middlewares

import (
	"net/http"
	"strings"

	"github.com/datto27/goecom/microservices/core/utils"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			utils.ApiError(c, http.StatusUnauthorized, "authorization header missing")
			c.Abort()
			return
		}

		// check for Bearer
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			utils.ApiError(c, http.StatusUnauthorized, "authorization header format must be Bearer {token}")
			c.Abort()
			return
		}

		token, err := utils.ParseJWT(tokenStr)

		if err != nil || !token.Valid {
			utils.ApiError(c, http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}

		c.Next()
	}
}
