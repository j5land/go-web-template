package router

import (
	"github.com/gin-gonic/gin"
	"go-web-template/pkg/gin/validator"
	"go-web-template/pkg/logger"
)

var routerEngine *gin.Engine

func init(){
	// 注册中文翻译器
	if err:= validator.InitTrans("zh");err != nil{
		logger.Fatal("init trans failed,err %v\n",err)
		return
	}
	routerEngine = gin.Default()
}

func GET(relativePath string, handlers gin.HandlerFunc) gin.IRoutes {
	return routerEngine.GET(relativePath, handlers)
}

func POST(relativePath string, handlers gin.HandlerFunc) gin.IRoutes {
	return routerEngine.POST(relativePath, handlers)
}

func DELETE(relativePath string, handlers gin.HandlerFunc) gin.IRoutes {
	return routerEngine.POST(relativePath, handlers)
}

func PATCH(relativePath string, handlers gin.HandlerFunc) gin.IRoutes {
	return routerEngine.PATCH(relativePath, handlers)
}

func OPTIONS(relativePath string, handlers gin.HandlerFunc) gin.IRoutes {
	return routerEngine.OPTIONS(relativePath, handlers)
}

func PUT(relativePath string, handlers gin.HandlerFunc) gin.IRoutes {
	return routerEngine.PUT(relativePath, handlers)
}

func HEAD(relativePath string, handlers gin.HandlerFunc) gin.IRoutes {
	return routerEngine.HEAD(relativePath, handlers)
}

func Run(addr string) error{
	return routerEngine.Run(addr)
}


