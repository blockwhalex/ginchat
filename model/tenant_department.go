package model

// 租户部门
type TenantDepartment struct {
	Model
	DepartmentName string `json:"department_name" gorm:"comment:部门名称"`
	Level          uint   `json:"level" gorm:"comment:部门级别"`
	Sort           uint   `json:"sort" gorm:"comment: 排序"`
	ParentId       uint   `json:"parent_id" gorm:"comment:上级部门ID"`
}
