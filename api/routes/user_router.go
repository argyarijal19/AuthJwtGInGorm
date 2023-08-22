package routes

import (
	"belajar-restapi/api/handlers"
	"belajar-restapi/api/middleware"
	DbConn "belajar-restapi/config"
	"belajar-restapi/repository/user"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	KoneksiDB := DbConn.Db_Mysql()

	data := user.UserTables(KoneksiDB)

	users := handlers.UserHandler{User: data}

	note := router.Group("/user")

	note.GET("/", middleware.AuthMiddleware(), users.ShowDataUSer)
	note.POST("/create_user", users.CreateUserNew)

	auth := router.Group("/auth")
	auth.POST("/", users.LoginUSer)
	auth.GET("/refreshtoken", users.RefreshToken)
}
