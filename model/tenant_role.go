package model

type TenantRole struct {
	Model
	RoleName string `json:"role_name" gorm:"租户角色"`
}
