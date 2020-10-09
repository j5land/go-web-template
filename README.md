# Go-Web-Template
**Go-Web-Template** 是一款基于Golang API微服务脚手架工程（模板工程）

- 版本最低支持，需求Go 1.13+
- 基于Go Modules

## 快速开始
快速构建你的基于RESTFul API微服务业务功能
```go
//1.绑定路由相对路径
router.POST("/login", func(c *gin.Context) {
    //2.创建参数绑定对象
    var login entity.Login
    //3.进行参数绑定与校验
    errString := validate.BindJSON(c, &login)
    //4.校验结果本地化返回
    if errString != "" {
        response.Fail(c, response.ResEntity{Message:errString})
        return
    }
    //5.通用成功响应
    response.Success(c, response.ResEntity{Data:gin.H{"user":login}})
})
```

## Go 应用程序项目目录规范
基于[Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_zh.md)

## 目前支持特性

基于此脚手架可以快速构建基于HTTP RESTFul的业务微服务

- 基于uber-zap 与 lumberjack.v2 日志模块
- 基于gin路由模块封装
- 基于gin数据绑定封装
- 基于validator支持中文校验封装
- 抽取API公用response简化代码
- 封装yaml.v2读取配置文件组件

## 未来路线图

未来计划加入以下组件特性，用于丰富脚手架工程

- 数据库ORM框架
- 定时调度任务
- 分布式缓存Redis
- 消息中间件Kafka
- Swagger
- JWT