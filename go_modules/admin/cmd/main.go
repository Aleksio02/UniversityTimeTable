package main

import (
	"admin/cmd/admin/config"
	"admin/cmd/admin/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	r := gin.New()

	v1 := r.Group("/admin")
	{
		//v1.Use(auth())
		v1.GET("/system/test", controller.SystemTest)
	}
	r.Run(fmt.Sprintf(":%v", config.Config.Application.Port))
}
