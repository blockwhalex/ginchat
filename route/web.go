package route

import (
	"ginchat/controller"
	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetWebGroupRoutes(router *gin.RouterGroup) {
	router.GET("/", controller.Index)
}
