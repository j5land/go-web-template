package main

import (
	"github.com/gin-gonic/gin"
	"go-web-template/internal/config"
	"go-web-template/internal/entity"
	"go-web-template/pkg/gin/response"
	"go-web-template/pkg/gin/router"
	validate "go-web-template/pkg/gin/validator"
	_ "go-web-template/pkg/requests"
)

func main() {

	router.GET("/hello", func(c *gin.Context) {
		response.Success(c, response.ResEntity{Message: "hello world"})
	})
	
	router.POST("/login", func(c *gin.Context) {
		var login entity.Login
		errString := validate.BindJSON(c, &login)
		if errString != "" {
			response.Fail(c, response.ResEntity{Message:errString})
			return
		}
		response.Success(c, response.ResEntity{Data:gin.H{"user":login}})
	})

	router.Run(":"+config.GetConf().Port)
}
