package controller

import (
	"fmt"
	"ginchat/util"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	fmt.Println(util.GetIpLocation("14.127.123.1"))

	ua := c.GetHeader("User-Agent")
	// "Mozilla/5.0 (Linux; U; Android 2.3.7; en-us; Nexus One Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"
	fmt.Println(util.ParseUserAgent(ua))
	//c.HTML(http.StatusOK, "index.html", gin.H{
	//	"visitor_id": util.GenerateUuid(),
	//})
}
