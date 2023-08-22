package middleware

import (
	"belajar-restapi/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization token is missing"})
			ctx.Abort()
		}
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret_key"), nil
		})

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse claims from token"})
			ctx.Abort()
			return
		}

		exparationTime := time.Unix(int64(claims["exp"].(float64)), 0)

		if time.Now().After(exparationTime) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Token has expired"})
			ctx.Abort()
			return
		}

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func RefreshTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshTokenString := c.GetHeader("Refresh-Token")
		if refreshTokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh token is missing"})
			c.Abort()
			return
		}

		refreshtoken, err := jwt.Parse(refreshTokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte("refresh_secret_key"), nil
		})

		if err != nil || !refreshtoken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid refresh token"})
			c.Abort()
			return
		}

		claims, ok := refreshtoken.Claims.(jwt.MapClaims)

		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse claims from refresh token"})
			c.Abort()
			return
		}

		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)

		if time.Now().After(expirationTime) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token has expired"})
			c.Abort()
			return
		}
		version := claims["version"].(float64) + 1
		user := &models.UserSimgoa{
			Username:    claims["username"].(string),
			Plasma:      claims["plasma"].(string),
			Description: claims["description"].(string),
			Password:    claims["password"].(string),
		}

		c.Set("user", user)
		c.Set("version", version) // Set the user in context for downstream handlers
		c.Next()
	}
}
