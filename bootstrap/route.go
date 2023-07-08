package bootstrap

import (
	"ginchat/common"
	"ginchat/route"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// 注册 api 分组路由
	apiGroup := router.Group("/api/v1")
	route.SetApiGroupRoutes(apiGroup)

	return router
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()
	err := r.Run(":" + common.App.Config.App.Port)
	if err != nil {
		return
	}
}
