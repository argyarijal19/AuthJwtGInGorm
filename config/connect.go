package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Db_Mysql() *gorm.DB {
	connectionString := os.Getenv("DATABASESTRINGCONNECT")

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Printf("Cant Connect to database : " + err.Error())
		panic(err)
	}

	return DB

}
