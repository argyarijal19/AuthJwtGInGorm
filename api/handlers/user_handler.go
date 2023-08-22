package handlers

import (
	"belajar-restapi/api/middleware"
	"belajar-restapi/models"
	"belajar-restapi/repository/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	User *user.UserTabel
}

// ShowDataUSer godoc
// @Summary melihat data User.
// @Description get data User.
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer Token"
// @Tags User Detail
// @Accept application/json
// @Produce json
// @Success 200 {object} models.ListDataUser
// @Router /user [get]
func (db *UserHandler) ShowDataUSer(c *gin.Context) {
	usr := new([]models.UserSimgoa)
	db.User.GetAllUser(usr)
	if len(*usr) == 0 {
		middleware.ErrorResponse(c, http.StatusNotFound, "Data Not Found", "")
		return
	}
	middleware.SuccessResponse(c, http.StatusOK, "success", usr)
}

// ShowDataUSer godoc
// @Summary create data User.
// @Description create data User.
// @Tags User Detail
// @Accept application/json
// @Param request body models.UserSimgoa true "Payload Body [RAW]"
// @Produce json
// @Success 200 {object} models.ListDataUser
// @Router /user/create_user [post]
func (db *UserHandler) CreateUserNew(c *gin.Context) {
	dataUser := new(models.UserSimgoa)

	err := c.ShouldBindJSON(&dataUser)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, "Register Failed", err.Error())
		return
	}

	err = db.User.CreateUser(dataUser)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	middleware.SuccessResponse(c, http.StatusOK, "Register Success", "success")

}
