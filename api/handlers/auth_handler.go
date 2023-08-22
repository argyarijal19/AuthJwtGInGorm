package handlers

import (
	"belajar-restapi/helper"
	"belajar-restapi/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ShowDataUSer godoc
// @Summary Login User.
// @Description Login User.
// @Tags Authentication
// @Accept application/json
// @Param request body models.LoginUser true "Payload Body [RAW]"
// @Produce json
// @Success 200 {object} models.LoginUser
// @Router /auth [post]
func (db *UserHandler) LoginUSer(c *gin.Context) {
	dataLogin := new(models.LoginUser)
	if err := c.ShouldBindJSON(&dataLogin); err != nil {
		c.JSON(http.StatusBadRequest, helper.ReturnData{
			Code:    400,
			Success: false,
			Status:  "Gagal Login",
			Data:    err.Error(),
		})
		return
	}
	user, err := db.User.GetUserByUsername(dataLogin.Username)

	if err != nil {
		c.JSON(http.StatusUnauthorized, helper.ReturnData{
			Code:    401,
			Success: false,
			Status:  "Invalid Username",
			Data:    nil,
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dataLogin.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, helper.ReturnData{
			Code:    401,
			Success: false,
			Status:  "Invalid Password",
			Data:    nil,
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"plasma":   user.Plasma,
		"exp":      time.Now().Add(time.Minute * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ReturnData{
			Code:    500,
			Success: false,
			Status:  "Internal Server Error",
			Data:    nil,
		})
		return
	}

	refreshTOkenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"plasma":   user.Plasma,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	}).SignedString([]byte("refresh_secret_key"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ReturnData{
			Code:    500,
			Success: false,
			Status:  "Internal Server Error",
			Data:    nil,
		})
		return
	}

	tokenGabung := map[string]interface{}{
		"access_token":  tokenString,
		"refresh_token": refreshTOkenString,
	}

	c.JSON(200, helper.ReturnData{
		Code:    200,
		Success: true,
		Status:  "Berhasil Login",
		Data:    tokenGabung,
	})
}

// RefreshToken godoc
// @Summary Refresh Access Token.
// @Description Refresh the access token using refresh token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Refresh-Token header string true "Refresh Token"
// @Success 200 {object} models.NewAccessToken
// @Router /auth/refreshtoken [get]
func (db *UserHandler) RefreshToken(c *gin.Context) {
	refreshTokenString := c.GetHeader("Refresh-Token")
	if refreshTokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh token is missing"})
		return
	}

	refreshtoken, err := jwt.Parse(refreshTokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte("refresh_secret_key"), nil
	})

	if err != nil || !refreshtoken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid refresh token"})
		return
	}

	claims, ok := refreshtoken.Claims.(jwt.MapClaims)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse claims from refresh token"})
		return
	}

	exparationTime := time.Unix(int64(claims["exp"].(float64)), 0)

	if time.Now().After(exparationTime) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token has expired"})
		c.Abort()
		return
	}

	newTokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": claims["username"],
		"plasma":   claims["plasma"],
		"exp":      time.Now().Add(time.Minute * 2).Unix(),
	})

	tokenStringNew, err := newTokenString.SignedString([]byte("secret_key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ReturnData{
			Code:    500,
			Success: false,
			Status:  "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": tokenStringNew,
	})
}
