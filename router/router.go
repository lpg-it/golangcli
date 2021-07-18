package router

import (
	"golangcli/controller"
	"golangcli/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 获取所有攻击方案
	v1.GET("/attack", controller.GetAttackHandler)
	// 增加攻击方案
	v1.POST("/attack", controller.CreateAttackHandler)
	// 修改攻击方案
	v1.PUT("/attack", controller.UpdateAttackHandler)
	// 删除攻击方案
	v1.DELETE("/attack", controller.DeleteAttackHandler)

	return r
}
