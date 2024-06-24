package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ko44d/hrmos-time-aggregator/pkg/controller"
)

func SetupRouter(eoc controller.EmployeeOvertimeController) *gin.Engine {
	r := gin.Default()

	// テンプレートのロード
	r.LoadHTMLGlob("templates/*")

	// ルーティングの設定
	r.GET("/api/v1/aggregate", eoc.Aggregate)

	return r
}
