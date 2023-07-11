package controller

import (
	"fmt"
	"ginchat/util"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	fmt.Println(util.GetIpLocation("14.127.123.1"))
	//c.HTML(http.StatusOK, "index.html", gin.H{
	//	"visitor_id": util.GenerateUuid(),
	//})
}
