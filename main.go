package main

import (
	"ginchat/bootstrap"
	"ginchat/common"
	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化配置
	bootstrap.InitConfig()
	// 初始化日志
	common.App.Log = bootstrap.InitializeLog()
	common.App.Log.Info("log init success!")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 启动服务
	err := r.Run(":" + common.App.Config.App.Port)
	if err != nil {
		return
	}
}
