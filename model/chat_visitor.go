package model

// 访客
type ChatVisitor struct {
	Model
	VisitorId  string `json:"visitor_id" gorm:"index:idx_visitor_id;comment:访客ID"`
	UserId     uint64 `json:"user_id" gorm:"comment:用户ID"`
	Nickname   string `json:"nickname" gorm:"comment:昵称"`
	Avatar     string `json:"avatar" gorm:"comment:头像"`
	SourceCode string `json:"source_code" gorm:"comment:来源"`
	Remark     string `json:"remark" gorm:"comment:备注"`
	Ip         string `json:"ip" gorm:"comment:IP"`
	Continent  string `json:"continent" gorm:"comment:洲"`
	Country    string `json:"country" gorm:"comment:国家"`
	TimeZone   string `json:"time_zone" gorm:"comment:时区"`
	IsoCode    string `json:"iso_code" gorm:"comment:国家简码"`
	Province   string `json:"province" gorm:"comment:省"`
	City       string `json:"city" gorm:"comment:市"`
	Latitude   string `json:"latitude" gorm:"纬度"`
	Longitude  string `json:"longitude" gorm:"经度"`
}
