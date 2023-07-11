package model

import "strconv"

// 租户用户
type TenantUser struct {
	Model
	UserId   string `json:"user_id" gorm:"index:idx_user_id;comment:用户ID"`
	NickName string `json:"nick_name" gorm:"comment:昵称"`
	Email    string `json:"email" gorm:"type:varchar(110);not_null;default '';unique_index;comment:邮箱"`
	Phone    string `json:"phone" gorm:"not null;index;comment:手机号"`
	Password string `json:"-" gorm:"not null;default:'';comment:用户密码"`
	Avatar   string `json:"avatar" gorm:"comment:头像"`
	Gender   uint   `json:"gender" gorm:"comment:性别"`
}

func (user TenantUser) GetUid() string {
	return strconv.Itoa(int(user.ID))
}
