package model

import (
	"time"
)

// 访客
type ChatVisitor struct {
	Model
	VisitorId string `json:"visitor_id" gorm:"index:idx_visitor_id;comment:访客ID"`
	UserId    string `json:"user_id" gorm:"comment:用户ID"`
	Nickname  string `json:"nickname" gorm:"comment:昵称"`
	Avatar    string `json:"avatar" gorm:"comment:头像"`
	Remark    string `json:"remark" gorm:"comment:备注"`
	Ip        string `json:"ip" gorm:"comment:IP"`
	Continent string `json:"continent" gorm:"comment:洲"`
	Country   string `json:"country" gorm:"comment:国家"`
	TimeZone  string `json:"time_zone" gorm:"comment:时区"`
	IsoCode   string `json:"iso_code" gorm:"comment:国家简码"`
	Province  string `json:"province" gorm:"comment:省"`
	City      string `json:"city" gorm:"comment:市"`
	Latitude  string `json:"latitude" gorm:"comment:纬度"`
	Longitude string `json:"longitude" gorm:"comment:经度"`
	Os        string `json:"os" gorm:"comment:操作系统"`
	Browser   string `json:"browser" gorm:"comment:浏览器"`
}

// 新增
func CreateVisitor(visitorId, userId, ip, continent, country, timezone, isoCode, province, city, latitude, longitude, os, browser string) uint {
	visitor := &ChatVisitor{
		VisitorId: visitorId,
		UserId:    userId,
		Ip:        ip,
		Continent: continent,
		Country:   country,
		TimeZone:  timezone,
		IsoCode:   isoCode,
		Province:  province,
		City:      city,
		Latitude:  latitude,
		Longitude: longitude,
		Os:        os,
		Browser:   browser,
	}
	visitor.UpdatedAt = time.Now()
	db.Create(visitor)
	return visitor.ID
}

// 查询
func GetVisitorByVisitorId(visitorId string) ChatVisitor {
	var visitor ChatVisitor
	db.Where("visitor_id = ?", visitor.VisitorId).First(&visitor)
	return visitor
}
