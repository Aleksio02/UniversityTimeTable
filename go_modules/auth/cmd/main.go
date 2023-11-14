package main

import (
	"auth/cmd/auth/config"
	"auth/cmd/auth/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	config.CreateDatabaseConnection()

	r := gin.New()

	v1 := r.Group("/auth")
	{
		//v1.Use(auth())
		v1.GET("/system/test", controller.SystemTest)
		v1.GET("/getSession", controller.GetSession)
		v1.GET("/createUser", controller.CreateUser)
		v1.POST("/sendAuthInfo", controller.MockTgBot)
	}
	r.Run(fmt.Sprintf(":%v", config.Config.Application.Port))
}
