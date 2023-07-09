package model

// 租户域名
type TenantDomain struct {
	Model
	Domain     string `json:"domain" gorm:"comment:域名"`
	ServerAddr string `json:"server_addr" gorm:"comment:服务器地址"`
}
