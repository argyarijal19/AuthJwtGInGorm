package middleware

import (
	"belajar-restapi/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateAccessToken(user *models.UserSimgoa) (string, error) {
	accesstoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":    user.Username,
		"plasma":      user.Plasma,
		"description": user.Description,
		"password":    user.Password,
		"exp":         time.Now().Add(time.Minute * 2).Unix(),
	})

	accessTokenString, err := accesstoken.SignedString([]byte("secret_key"))
	if err != nil {
		return "", nil
	}

	return accessTokenString, nil
}

func GenerateRefreshToken(user *models.UserSimgoa) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":    user.Username,
		"plasma":      user.Plasma,
		"description": user.Description,
		"password":    user.Password,
		"exp":         time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	refreshTokenKey := []byte("refresh_secret_key")
	refreshTokenString, err := refreshToken.SignedString(refreshTokenKey)
	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}
