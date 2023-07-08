package main

import (
	"ginchat/bootstrap"
	"ginchat/common"
)

func main() {

	// 初始化配置
	bootstrap.InitConfig()
	// 初始化日志
	common.App.Log = bootstrap.InitializeLog()
	common.App.Log.Info("log init success!")

	// 初始化数据库
	common.App.DB = bootstrap.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if common.App.DB != nil {
			db, _ := common.App.DB.DB()
			err := db.Close()
			if err != nil {
				return
			}
		}
	}()

	// 启动服务器
	bootstrap.RunServer()
}
