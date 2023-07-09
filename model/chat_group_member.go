package model

// 聊天群成员
type ChatGroupMember struct {
	Model
	ChatGroupId   string `json:"chat_group_id" gorm:"群组ID"`
	MemberId      string `json:"member_id" gorm:"comment:成员ID"`
	IsGroupLeader string `json:"is_group_leader" gorm:"comment:是否是群主"`
}
