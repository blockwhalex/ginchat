package controller

import (
	"ginchat/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"visitor_id": util.GenerateUuid(),
	})
}
