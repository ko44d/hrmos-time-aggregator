package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ko44d/hrmos-time-aggregator/pkg/controller"
)

func SetupRouter(tpc controller.TopPageController, eoc controller.EmployeeOvertimeController) *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", tpc.ShowForm)
	r.POST("/set_api_key", tpc.SetAPIKey)
	r.GET("/aggregate", eoc.Aggregate)

	return r
}
