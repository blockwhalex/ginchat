package model

// 聊天服务评价
type ChatServiceComment struct {
	Model
	VisitorId      string `json:"visitor_id" gorm:"comment:访客ID"`
	ChatId         uint64 `json:"chat_id" gorm:"comment:会话ID"`
	UserId         uint64 `json:"user_id" gorm:"comment:客服ID"`
	TerminalSource string `json:"terminal_source" gorm:"comment:终端来源"`
	Tag            string `json:"tag" gorm:"comment:评价标签"`
	Score          uint   `json:"score" gorm:"comment:五星评分"`
	Remark         string `json:"remark" gorm:"comment:备注"`
	CommentType    string `json:"comment_type" gorm:"comment:评价类型,主动-active 邀请-invite"`
	IsSolved       uint   `json:"is_solved" gorm:"comment:是否解决"`
}
