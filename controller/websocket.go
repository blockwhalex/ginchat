package controller

import (
	"ginchat/service/websocket"
	"github.com/gin-gonic/gin"
)

func Websocket(c *gin.Context) {
	websocket.MsgHandler(c)
	return
}
