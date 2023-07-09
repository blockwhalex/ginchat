package model

import "time"

// 聊天消息
type ChatMessage struct {
	Model
	MsgType      string    `json:"msg_type" gorm:"comment:消息类型"`
	ContentType  string    `json:"content_type" gorm:"comment:内容类型"`
	Content      string    `json:"content" gorm:"comment:内容"`
	SenderId     string    `json:"sender_id" gorm:"comment:发送者ID"`
	SenderType   string    `json:"sender_type" gorm:"comment:发送放类型,访客-visitor 客服-customer_service 机器人-robot"`
	ReceiverId   string    `json:"receiver_id" gorm:"comment:接收人ID"`
	ReceiveType  string    `json:"receive_type" gorm:"comment:接收方类型"`
	SendTime     time.Time `json:"send_time" gorm:"comment:发送时间"`
	SourceCode   string    `json:"source_code" gorm:"comment:来源"`
	IsOfflineMsg uint      `json:"is_offline_msg" gorm:"comment:是否是离线消息"`
}
