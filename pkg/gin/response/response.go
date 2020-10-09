package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 自定义响应数据结构体
type ResEntity struct {
	// HTTP 状态码 & 自定义状态码
	Code    int
	// 输出消息
	Message string
	// 输出自定义数据
	Data  	gin.H
}

// 通用响应成功处理结果信息拼装
func Success(c *gin.Context, entity ResEntity){
	code := http.StatusOK
	message := http.StatusText(code)
	if entity.Code != 0 {
		code = entity.Code
	}
	if entity.Message != "" {
		message = entity.Message
	}
	res := assembly(code, message, entity.Data)
	c.JSON(code, res)
}

// 通用失败成功处理结果信息拼装
func Fail(c *gin.Context, entity ResEntity){
	code := http.StatusBadRequest
	message := http.StatusText(code)
	if entity.Code != 0 {
		code = entity.Code
	}
	if entity.Message != "" {
		message = entity.Message
	}
	res := assembly(code, message, entity.Data)
	c.JSON(code, res)
}

// 组装返回结果信息
func assembly(code int, message string, data gin.H) gin.H {
	response := gin.H{
		"code":    code,
		"message": message,
	}
	response["data"] = data
	/*for index, value := range data {
		response[index] = value
	}*/
	return response
}
