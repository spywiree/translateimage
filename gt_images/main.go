package main

import (
	"github.com/gin-gonic/gin"
)

func main8() {
	r := gin.Default()

	// 注册POST路由到handlePost函数
	r.POST("/translation", translation)
	r.POST("/aliyunoss", aliyunoss)

	// 启动服务器并监听端口
	r.Run(":8090")
}
