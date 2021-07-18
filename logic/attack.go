package logic

import (
	"golangcli/dao/mysql"
	"golangcli/models"
)

// GetAttackList 获取所有攻击方案
func GetAttackList() (data []*models.AttackPlan, err error) {
	// 查询数据库,查到所有攻击方案并且返回
	data, err = mysql.GetAttackList()
	return
}

// CreateAttackPlan 增加攻击方案
func CreateAttackPlan(attackPlan *models.AttackPlan) (err error) {
	// 保存到数据库
	err = mysql.CreateAttackPlan(attackPlan)
	if err != nil {
		return err
	}
	return err
}

// DeleteAttackPlan 删除攻击方案
func DeleteAttackPlan(attackPlan *models.AttackPlan) (err error) {
	// 保存到数据库
	err = mysql.DeleteAttackPlan(attackPlan)
	if err != nil {
		return err
	}
	return err
}

// UpdateAttackPlan 修改攻击方案
func UpdateAttackPlan(attackPlan *models.AttackPlan) (err error) {
	// 保存到数据库
	err = mysql.UpdateAttackPlan(attackPlan)
	if err != nil {
		return err
	}
	return err
}
