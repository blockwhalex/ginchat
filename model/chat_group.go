package model

// 聊天群
type ChatGroup struct {
	Model
	Name         string `json:"name" gorm:"not null;comment:群组名称"`
	Icon         string `json:"icon" gorm:"群图标"`
	Description  string `json:"description" gorm:"not null;comment:群组描述"`
	Type         string `json:"type" gorm:"not null;comment:类型"` // public-公开群 private-私有群
	NeedApproval string `json:"need_approval" gorm:"comment:是否需要审批(群主或管理员)"`
	MaxLimit     int    `json:"max_limit" gorm:"comment:成员上限"`
	CurrentNum   int    `json:"current_num" gorm:"comment:当前人数"`
	IsMuted      string `json:"has_muted" gorm:" comment:是否全员禁言"`
}
