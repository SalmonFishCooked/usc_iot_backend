package main

import (
	"github.com/gin-gonic/gin"
	"usc_iot_backend/common"
	"usc_iot_backend/controller"
)

// CollectRoute 创建路由组
func CollectRoute(r *gin.Engine) *gin.Engine {
	//跨域中间件
	r.Use(common.Cors())

	//API总前缀路由组
	apiGroup := r.Group("/api")

	//设备路由组
	deviceGroup := apiGroup.Group("/device")
	deviceGroup.POST("/info", controller.GetDeviceInfo)

	//传感器路由组
	sensorGroup := apiGroup.Group("/sensor")
	sensorGroup.POST("/info", controller.GetSensorInfo)
	sensorGroup.POST("/create", controller.CreateSensor)
	sensorGroup.POST("/delete", controller.DeleteSensor)

	//历史传感数据路由组
	historyGroup := apiGroup.Group("/history")
	historyGroup.POST("/info", controller.GetHistoryInfo)
	//historyGroup.POST("/create", controller.CreateSensor)
	//historyGroup.POST("/delete", controller.DeleteSensor)

	return r
}
