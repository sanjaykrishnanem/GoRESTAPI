package main

import (
	"CRUDTEST/Config"
	"CRUDTEST/Models"
	"CRUDTEST/Routes"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//running
	r.Run()
}
