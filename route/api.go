package route

import (
	"ginchat/common"
	"ginchat/controller"
	"ginchat/middleware"
	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/register", controller.Register)
	router.POST("/auth/login", controller.Login)

	authRouter := router.Group("").Use(middleware.JWTAuth(common.AppGuardName))
	{
		authRouter.POST("/auth/info", controller.Info)
	}
}
