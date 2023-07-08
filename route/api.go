package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})
}
