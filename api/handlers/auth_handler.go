package handlers

import (
	"belajar-restapi/api/middleware"
	"belajar-restapi/helper"
	"belajar-restapi/models"
	"net/http"

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

	accessToken, err := middleware.GenerateAccessToken(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ReturnData{
			Code:    500,
			Success: false,
			Status:  "Internal Server Error",
			Data:    err.Error(),
		})
		return
	}

	refreshToken, err := middleware.GenerateRefreshToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ReturnData{
			Code:    500,
			Success: false,
			Status:  "Internal Server Error",
			Data:    err.Error(),
		})
		return
	}

	tokenGabung := map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
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
	user, _ := c.Get("user")
	newTokenString, err := middleware.GenerateAccessToken(user.(*models.UserSimgoa))

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
		"access_token": newTokenString,
	})
}
