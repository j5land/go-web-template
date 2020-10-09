package main

import (
	"github.com/gin-gonic/gin"
	"golang-web-template/internal/config"
	"golang-web-template/internal/entity"
	"golang-web-template/pkg/gin/response"
	"golang-web-template/pkg/gin/router"
	validate "golang-web-template/pkg/gin/validator"
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
