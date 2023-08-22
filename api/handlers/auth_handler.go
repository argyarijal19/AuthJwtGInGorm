package handlers

import (
	"belajar-restapi/api/middleware"
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
		middleware.ErrorResponse(c, http.StatusBadRequest, "Gagal Login", err.Error())
		return
	}
	user, err := db.User.GetUserByUsername(dataLogin.Username)

	if err != nil {
		middleware.ErrorResponse(c, http.StatusUnauthorized, "Invalid Username", err.Error())
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dataLogin.Password)); err != nil {
		middleware.ErrorResponse(c, http.StatusUnauthorized, "Invalid Password", err.Error())
		return
	}

	accessToken, err := middleware.GenerateAccessToken(user)

	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	refreshToken, err := middleware.GenerateRefreshToken(user)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	tokenGabung := map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	middleware.SuccessResponse(c, http.StatusOK, "success", tokenGabung)
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
		middleware.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	middleware.SuccessResponse(c, http.StatusOK, "success", gin.H{"access_token": newTokenString})
}
