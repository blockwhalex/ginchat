package util

import (
	"github.com/mssola/useragent"
)

func ParseUserAgent(userAgent string) map[string]interface{} {
	ua := useragent.New(userAgent)
	uaInfo := make(map[string]interface{})
	uaInfo["is_mobile"] = ua.Mobile()        // 是否是移动端, bool
	uaInfo["is_robot"] = ua.Bot()            // 是否是机器人, bool
	uaInfo["mozilla_version"] = ua.Mozilla() // 火狐浏览器版本号
	uaInfo["device"] = ua.Model()            // 设备型号
	uaInfo["platform"] = ua.Platform()       // 平台
	uaInfo["os"] = ua.OS()                   // 操作系统
	enginName, engineVersion := ua.Engine()
	uaInfo["engine"] = enginName + "+" + engineVersion // 浏览器引擎及版本
	name, version := ua.Browser()
	uaInfo["browser"] = name + "-" + version // 浏览器及版本
	return uaInfo
}
