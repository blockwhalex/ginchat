package model

// 访客
type ChatVisitor struct {
	Model
	UserId     uint64 `json:"user_id" gorm:"comment:用户ID"`
	Name       string `json:"name" gorm:"comment:昵称"`
	Nickname   string `json:"nickname" gorm:"comment:昵称"`
	Avatar     string `json:"avatar" gorm:"comment:头像"`
	SourceCode string `json:"source_code" gorm:"comment:来源"`
	Remark     string `json:"remark" gorm:"comment:备注"`
	Ip         string `json:"ip" gorm:"comment:IP"`
	Province   string `json:"province" gorm:"comment:省"`
	City       string `json:"city" gorm:"comment:市"`
	Area       string `json:"area" gorm:"comment:区/县"`
}
