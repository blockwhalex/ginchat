package model

// 租户部门成员
type TenantDepartMem struct {
	DepartmentId string `json:"department_id" gorm:"comment:部门ID"`
	UserId       string `json:"user_id" gorm:"comment:用户ID"`
}
