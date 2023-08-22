package main

import (
	"github.com/joho/godotenv"

	setuproutes "belajar-restapi/api/routes/setupRoutes"
	_ "belajar-restapi/docs"
)

// @title Restful-API-Learn
// @description API Untuk belajar membuat endpoint dan membuat authetication menggunaka Jason Web Token
// @version 2.0

// @host localhost
// @BasePath /
// @schemes https
// @Failure      400  {object}	ReturnData
// @Failure      404  {object}  ReturnData
// @Failure      500  {object}  ReturnData

func main() {
	_ = godotenv.Load(".env")
	router := setuproutes.SetupAllRoutes()

	// Jalankan server pada port 8080
	router.Run(":8080")
}
