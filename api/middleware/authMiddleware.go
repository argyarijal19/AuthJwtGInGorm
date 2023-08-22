package middleware

import (
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
