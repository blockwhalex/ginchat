package controller

import (
	"ginchat/service"
	"github.com/gin-gonic/gin"
)

func Websocket(c *gin.Context) {
	service.MsgHandler(c)
	return
}
