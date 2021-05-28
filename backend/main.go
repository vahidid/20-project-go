package main

import (
	"fmt"
	"vahidid/20-project-go/Config"
	"vahidid/20-project-go/Models"
	"vahidid/20-project-go/Routes"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {
	dsn := "root:V@#!d1377@tcp(127.0.0.1:3306)/20-project?charset=utf8&parseTime=True&loc=Local"
	Config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("status: ", err)
	}

	// defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Task{})

	r := Routes.SetupRouter()

	r.Run("0.0.0.0:5000")
}
