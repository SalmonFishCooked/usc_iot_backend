package main

import (
	"github.com/gin-gonic/gin"
	"usc_iot_backend/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/device", controller.GetDeviceInfo)

	return r
}
