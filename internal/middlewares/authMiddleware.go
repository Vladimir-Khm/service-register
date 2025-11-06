package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"service-register/internal/config"
	"service-register/internal/models"
)

type AuthMiddleware struct {
	config *config.Config
}

func (middleware *AuthMiddleware) GetEnv() *config.Config {
	return middleware.config
}

func CreateAuthMiddleware(config *config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		config: config,
	}
}

func (middleware *AuthMiddleware) Use() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &models.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(middleware.GetEnv().PayloadSignatureKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*models.JwtCustomClaims); ok && token.Valid {
			c.Set("userID", claims.UserTelegramID)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		}
	}
}
