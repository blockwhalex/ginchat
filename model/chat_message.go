package model

import "time"

// 聊天消息
type ChatMessage struct {
	Model
	MsgType     string `json:"msg_type" gorm:"comment:消息类型"`
	ContentType string `json:"content_type" gorm:"comment:内容类型"`
	Content     string `json:"content" gorm:"comment:内容"`
	SenderId    string `json:"sender_id" gorm:"index:idx_sender_id;comment:发送者ID"`
	SenderType  string `json:"sender_type" gorm:"comment:发送放类型,访客-visitor 客服-customer_service 机器人-robot"`
	ReceiverId  string `json:"receiver_id" gorm:"index:idx_receiver_id;comment:接收人ID"`
	ReceiveType string `json:"receive_type" gorm:"comment:接收方类型"`
}

// 创建
func CreateMessage(msgType, contentType, content, senderId, sendType, receiverId, receiveType string) uint {
	message := ChatMessage{
		MsgType:     msgType,
		ContentType: contentType,
		Content:     content,
		SenderId:    senderId,
		SenderType:  sendType,
		ReceiverId:  receiverId,
		ReceiveType: receiveType,
	}
	message.UpdatedAt = time.Now()
	db.Create(message)
	return message.ID
}
