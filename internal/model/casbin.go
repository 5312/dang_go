package model

import (
	. "dang_go/internal/enforcer"
	"gorm.io/gorm"
)

type CasbinModel struct {
	PType  string `json:"p_type" gorm:"column:p_type" description:"策略类型"`
	RoleId string `json:"role_id" gorm:"column:v0" description:"角色ID"`
	Path   string `json:"path" gorm:"column:v1" description:"api路径"`
	Method string `json:"method" gorm:"column:v2" description:"访问方法"`
}

func (c *CasbinModel) TableName() string {
	return "casbin_rule"
}

func (c *CasbinModel) Create() error {
	e := EnforcerTool()
	if success, errors := e.AddPolicy(c.RoleId, c.Path, c.Method); success == false {
		return errors
	}
	return nil
}

func (c *CasbinModel) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(c).Where("v1 = ? AND v2 = ?", c.Path, c.Method).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (c *CasbinModel) List() [][]string {
	e := EnforcerTool()
	policy := e.GetFilteredPolicy(0, c.RoleId)
	return policy
}
