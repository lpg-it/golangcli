package mysql

import (
	"database/sql"
	"golangcli/models"

	"go.uber.org/zap"
)

// GetAttackList 查询所有攻击方案
func GetAttackList() (attackPlanList []*models.AttackPlan, err error) {
	sqlStr := "select attack_plan_id, attack_plan_name, attack_plan_type, attack_plan_content, attack_plan_use_count, attack_plan_use_time, attack_plan_finish_time, attack_plan_status, attack_plan_remark from attack_plan"
	if err := db.Select(&attackPlanList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no attack plan in db")
			err = nil
		}
	}

	return
}

// CreateAttackPlan 增加攻击方案
func CreateAttackPlan(attackPlan *models.AttackPlan) (err error) {
	sqlStr := `insert into attack_plan(attack_plan_id, attack_plan_name, attack_plan_type, attack_plan_content, attack_plan_use_count, attack_plan_use_time, attack_plan_finish_time, attack_plan_status, attack_plan_remark) values(?,?,?,?,?,?,?,?,?)`
	_, err = db.Exec(sqlStr, attackPlan.ID, attackPlan.Name, attackPlan.Type, attackPlan.Content, attackPlan.UseCount, attackPlan.UseTime, attackPlan.FinishTime, attackPlan.Status, attackPlan.Remark)

	return
}

// DeleteAttackPlan 删除攻击方案
func DeleteAttackPlan(attackPlan *models.AttackPlan) (err error) {
	sqlStr := "DELETE FROM attack_plan WHERE attack_plan_id = ?"
	_, err = db.Exec(sqlStr, attackPlan.ID)

	return
}

// UpdateAttackPlan 修改攻击方案
func UpdateAttackPlan(attackPlan *models.AttackPlan) (err error) {
	sqlStr := "UPDATE attack_plan SET attack_plan_name = ?, attack_plan_type = ?, attack_plan_content = ?, attack_plan_use_count = ?, attack_plan_use_time = ?, attack_plan_finish_time = ?, attack_plan_status = ?, attack_plan_remark = ? WHERE attack_plan_id = ?"
	_, err = db.Exec(sqlStr, attackPlan.Name, attackPlan.Type, attackPlan.Content, attackPlan.UseCount, attackPlan.UseTime, attackPlan.FinishTime, attackPlan.Status, attackPlan.Remark, attackPlan.ID)

	return
}
