package model

// 租户
type Tenant struct {
	Model
	CompanyName   string  `json:"company_name" gorm:"comment:公司名称"`
	Contact       string  `json:"contact" gorm:"comment:联系人"`
	Phone         string  `json:"phone" gorm:"comment:电话"`
	Email         string  `json:"email" gorm:"comment:邮箱"`
	DepartmentNum uint    `json:"department_num" gorm:"comment:部门数"`
	UserNum       uint    `json:"user_num" gorm:"comment:用户数"`
	Domain        string  `json:"domain" gorm:"comment:域名"`
	CompanySize   string  `json:"company_size" gorm:"comment:公司规模"`
	Balance       float32 `gorm:"type:decimal(10,2);comment:余额"`
}
