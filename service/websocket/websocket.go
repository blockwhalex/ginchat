package websocket

import "github.com/gin-gonic/gin"

func MsgHandler(c *gin.Context) {
	hub := newHub()
	go hub.run()
	serveWs(hub, c.Writer, c.Request)
}
