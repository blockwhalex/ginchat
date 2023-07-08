package middleware

import (
	"ginchat/common"
	"ginchat/common/response"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func CustomRecovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(
		&lumberjack.Logger{
			Filename:   common.App.Config.Log.RootDir + "/" + common.App.Config.Log.Filename,
			MaxSize:    common.App.Config.Log.MaxSize,
			MaxBackups: common.App.Config.Log.MaxBackups,
			MaxAge:     common.App.Config.Log.MaxAge,
			Compress:   common.App.Config.Log.Compress,
		},
		response.ServerError)
}
