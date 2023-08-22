package handlers

import (
	"belajar-restapi/helper"
	"belajar-restapi/models"
	"belajar-restapi/repository/user"

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
		c.JSON(404, helper.ReturnData{
			Code:    404,
			Success: false,
			Status:  "Data Not Found!!",
			Data:    nil,
		})
	}
	c.JSON(200, helper.ReturnData{
		Code:    200,
		Success: true,
		Status:  "Berhasil Get Data",
		Data:    usr,
	})
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
		c.JSON(400, helper.ReturnData{
			Code:    400,
			Success: false,
			Status:  "Bad Request",
			Data:    err.Error(),
		})
		return
	}

	err = db.User.CreateUser(dataUser)
	if err != nil {
		c.JSON(500, helper.ReturnData{
			Code:    500,
			Success: false,
			Status:  "Internal Server Error",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(200, helper.ReturnData{
		Code:    200,
		Success: true,
		Status:  "Berhasil post user",
		Data:    dataUser,
	})

}
