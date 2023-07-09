package model

import "time"

// 聊天
type Chat struct {
	Model
	VisitorId        string    `json:"visitor_id" gorm:"comment:访客ID"`
	UserId           string    `json:"user_id" gorm:"comment:客服ID"`
	StartTime        time.Time `json:"start_time" gorm:"comment:开始时间"`
	EndTime          time.Time `json:"end_time" gorm:"comment:结束时间"`
	SessionDuration  uint64    `json:"session_duration" gorm:"comment:会话时长，单位毫秒"`
	QueueDuration    uint64    `json:"queue_duration" gorm:"comment:排队时长"`
	QueueStatus      string    `json:"queue_status" gorm:"comment:排队状态 未接通-nit_connected 接通-connected 离开-leave"`
	RobotMsgNum      uint      `json:"robot_msg_num" gorm:"comment:咨询机器人消息数"`
	RobotReplyNum    uint      `json:"robot_reply_num" gorm:"comment:机器人回复数"`
	CustomerMsgNum   uint      `json:"customer_msg_num" gorm:"comment:客服消息数"`
	CustomerReplyNum uint      `json:"customer_reply_num" gorm:"comment:人工回复数"`
	TransferStatus   string    `json:"transfer_status" gorm:"comment:转人工状态 成功-success  失败-fail"`
	Ip               string    `json:"ip" gorm:"comment:IP"`
	Os               string    `json:"os" gorm:"comment:操作系统,Windows Linux macOS Android iOS"`
	OfflineType      string    `json:"offline_type" gorm:"comment:离线方式 1 客服离线，2 客户被客服移除 3 客户被客服加入黑名单 4 客户会话超时 5 客户关闭了聊天页面 6 客户打开新的页面"`
	SearchSource     string    `json:"search_source" gorm:"comment:搜索来源 1、百度自然搜索；2、百度付费搜索；3、360搜索；4、sougou;5、神马；6、必应；7、谷歌；8、其他搜索引擎 ；9、直接访问； 10、外部链接；11、百度未知访问"` //
	SearchWords      string    `json:"search_words" gorm:"comment:搜索词"`
	LandPageUrl      string    `json:"land_page_url" gorm:"comment:落地页URL"`
	LandPageTitle    string    `json:"land_page_title" gorm:"comment:落地页标题"`
}
