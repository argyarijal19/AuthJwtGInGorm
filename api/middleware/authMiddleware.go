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
			ErrorResponse(ctx, http.StatusNotFound, "Invalid Key", "")
			ctx.Abort()
		}
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret_key"), nil
		})
		if err != nil {
			ErrorResponse(ctx, http.StatusForbidden, "Invalid token", err.Error())
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ErrorResponse(ctx, http.StatusForbidden, "Invalid token", err.Error())
			ctx.Abort()
			return
		}

		exparationTime := time.Unix(int64(claims["exp"].(float64)), 0)

		if time.Now().After(exparationTime) {
			ErrorResponse(ctx, http.StatusForbidden, "Token Has Expired", "")
			ctx.Abort()
			return
		}

		if err != nil || !token.Valid {
			ErrorResponse(ctx, http.StatusForbidden, "Invalid Key", err.Error())
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
			ErrorResponse(c, http.StatusNotFound, "Token Not Defined", "")
			c.Abort()
			return
		}

		refreshtoken, err := jwt.Parse(refreshTokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte("refresh_secret_key"), nil
		})

		if err != nil || !refreshtoken.Valid {
			ErrorResponse(c, http.StatusForbidden, "Invalid refresh token", err.Error())
			c.Abort()
			return
		}

		claims, ok := refreshtoken.Claims.(jwt.MapClaims)

		if !ok {
			ErrorResponse(c, http.StatusForbidden, "Failed to parse claims from refresh token", "")
			c.Abort()
			return
		}

		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)

		if time.Now().After(expirationTime) {
			ErrorResponse(c, http.StatusUnauthorized, "token has expired", "")
			c.Abort()
			return
		}
		user := &models.UserSimgoa{
			Username:    claims["username"].(string),
			Plasma:      claims["plasma"].(string),
			Description: claims["description"].(string),
			Password:    claims["password"].(string),
		}

		c.Set("user", user)
		c.Next()
	}
}
