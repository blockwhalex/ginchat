package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primaryKey;comment:ID"`
	TenantId  uint           `json:"tenant_id" gorm:"comment:租户ID"`
	AppId     uint           `json:"app_id" gorm:"comment:应用ID"`
	CreatedAt time.Time      `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;comment:删除时间"`
}
