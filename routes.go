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
	deviceGroup.POST("/create", controller.CreateDevice)
	deviceGroup.POST("/delete", controller.DeleteDevice)

	//传感器路由组
	sensorGroup := apiGroup.Group("/sensor")
	sensorGroup.POST("/info", controller.GetSensorInfo)
	sensorGroup.POST("/create", controller.CreateSensor)
	sensorGroup.POST("/delete", controller.DeleteSensor)

	//执行器路由组
	ActuatorGroup := apiGroup.Group("/actuator")
	ActuatorGroup.POST("/info", controller.GetActuatorInfo)
	ActuatorGroup.POST("/create", controller.CreateActuator)
	ActuatorGroup.POST("/delete", controller.DeleteActuator)

	//历史传感数据路由组
	historyGroup := apiGroup.Group("/history")
	historyGroup.POST("/info", controller.GetHistoryInfo)
	historyGroup.POST("/create", controller.CreateHistory)
	historyGroup.POST("/delete", controller.DeleteHistory)

	//电池数据路由组
	batteryGroup := apiGroup.Group("/battery")
	batteryGroup.POST("/info", controller.GetBatteryInfo)
	//batteryGroup.POST("/create", controller.CreateBattery)
	//batteryGroup.POST("/delete", controller.DeleteBattery)

	return r
}
