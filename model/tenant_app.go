package model

// 租户应用
type TenantApp struct {
	Model
	AppName string `json:"app_name" gorm:"not null;comment:应用名称"`
	AppKey  string `json:"app_key" gorm:"not null;comment:AppKey"`
}
