package controller

import (
	"golangcli/logic"
	"golangcli/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func GetAttackHandler(c *gin.Context) {
	// 查询到所有的攻击计划, 以切片形式返回
	data, err := logic.GetAttackList()

	if err != nil {
		zap.L().Error("logic.GetAttackList failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, data)
}

func CreateAttackHandler(c *gin.Context) {
	// 增加攻击方案
	// 获取参数
	attackPlan := new(models.AttackPlan)
	if err := c.ShouldBindJSON(attackPlan); err != nil {
		zap.L().Debug("c.ShouldBindJSON(attackPlan) err", zap.Any("err", err))
		zap.L().Error("CreateAttackPlan with invalid param failed")
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err := logic.CreateAttackPlan(attackPlan); err != nil {
		zap.L().Error("logic.CreateAttackPlan failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	ResponseSuccess(c, nil)
}

func UpdateAttackHandler(c *gin.Context) {
	// 修改攻击方案
	// 获取参数
	attackPlan := new(models.AttackPlan)
	if err := c.ShouldBindJSON(attackPlan); err != nil {
		zap.L().Debug("c.ShouldBindJSON(attackPlan) err", zap.Any("err", err))
		zap.L().Error("CreateAttackPlan with invalid param failed")
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err := logic.UpdateAttackPlan(attackPlan); err != nil {
		zap.L().Error("logic.UpdateAttackPlan failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	ResponseSuccess(c, nil)
}

func DeleteAttackHandler(c *gin.Context) {
	// 删除攻击方案
	// 获取参数
	attackPlan := new(models.AttackPlan)
	if err := c.ShouldBindJSON(attackPlan); err != nil {
		zap.L().Debug("c.ShouldBindJSON(attackPlan) err", zap.Any("err", err))
		zap.L().Error("CreateAttackPlan with invalid param failed")
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err := logic.DeleteAttackPlan(attackPlan); err != nil {
		zap.L().Error("logic.UpdateAttackPlan failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	ResponseSuccess(c, nil)
}
